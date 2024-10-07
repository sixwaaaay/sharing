/*
 * Copyright (c) 2023-2024 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.sixwaaaay.sharingcomment.service;


import io.sixwaaaay.sharingcomment.client.UserClient;
import io.sixwaaaay.sharingcomment.client.VoteClient;
import io.sixwaaaay.sharingcomment.domain.*;
import io.sixwaaaay.sharingcomment.repository.CommentRepository;
import io.sixwaaaay.sharingcomment.request.Principal;
import io.sixwaaaay.sharingcomment.util.DbContext;
import io.sixwaaaay.sharingcomment.util.DbContextEnum;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Limit;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import static java.util.function.Function.identity;

@Service
public class CommentService {
    private final CommentRepository commentRepo;

    private final VoteClient voteClient;

    private final UserClient userClient;

    private final boolean enableVote;
    private final boolean enableUser;

    public CommentService(CommentRepository commentRepo, VoteClient voteClient, UserClient userClient, @Value("${service.vote.enabled}") boolean enableVote, @Value("${service.user.enabled}") boolean enableUser) {
        this.commentRepo = commentRepo;
        this.voteClient = voteClient;
        this.userClient = userClient;
        this.enableVote = enableVote;
        this.enableUser = enableUser;
    }

    /**
     * This method is used to get the main comment list.
     * It first retrieves the count of comments belonging to the same entity.
     * If the count is zero, it returns an empty list of comments.
     * Otherwise, it retrieves the main comments that belong to the same entity and have an id less than the provided id.
     * For each main comment that has reply comments, it retrieves the latest 2 reply comments.
     * Then, it composes the comments by filling in the user info and vote status.
     * Finally, it sets the comments and the next page id to the result.
     *
     * @param belongTo the id of the object which the comments belong to
     * @param id       the id of the earliest comment in the previous page
     * @param size     the number of comments to be retrieved
     * @param userId   the id of the user who is requesting
     * @return a CommentResult object that contains the total count of comments, the comments for the current page, and the id of the next page
     */
    public CommentResult getMainCommentList(Long belongTo, Long id, Integer size, Long userId) {
        DbContext.set(DbContextEnum.READ);  /* set read context */

        var limit = Limit.of(size);
        var mainComments = commentRepo.findByBelongToAndIdLessThanAndReplyToNullOrderByIdDesc(belongTo, id, limit);
        /* for each main comment which has reply comments, get the latest 2 reply comments */
        mainComments.stream().filter(comment -> comment.getReplyCount() != 0).forEach(comment -> {
            var replyComments = commentRepo.findByBelongToAndReplyToAndIdGreaterThanOrderByIdAsc(belongTo, comment.getId(), 0L, Limit.of(2));
            comment.setReplyComments(replyComments);
        });

        composeComment(mainComments, userId);

        var result = new CommentResult();

        result.setComments(mainComments);
        if (mainComments.size() == limit.max()) {
            result.setNextPage(mainComments.getLast().getId());
        }
        return result;
    }

    /**
     * This method is used to get the reply comment list.
     * It first retrieves the reply comments that belong to the same comment and have an id greater than the provided id.
     * Then, it composes the comments by filling in the user info and vote status.
     * Finally, it sets the comments and the next page id to the result.
     *
     * @param belongTo the id of the object which the comments belong to
     * @param replyTo  the id of the comment which the comments are replies to
     * @param id       the id of the latest comment in the previous page
     * @param size     the number of comments to be retrieved
     * @param userId   the id of the user who is requesting
     * @return a ReplyResult object that contains the comments for the current page and the id of the next page
     */
    public ReplyResult getReplyCommentList(Long belongTo, Long replyTo, Long id, Integer size, Long userId) {
        DbContext.set(DbContextEnum.READ); /* set read context */
        var limit = Limit.of(size);
        var comments = commentRepo.findByBelongToAndReplyToAndIdGreaterThanOrderByIdAsc(belongTo, replyTo, id, limit);

        /* sort by id asc */
        comments.sort(Comparator.comparingLong(Comment::getId));
        composeComment(comments, userId);

        var result = new ReplyResult();
        result.setComments(comments);

        if (comments.size() == limit.max()) {
            result.setNextPage(comments.getLast().getId());
        }
        return result;
    }

    /**
     * This method is used to create a new comment in the system.
     * Then, it increases the count of comments belonging to the same entity.
     * Finally, it composes the comment by filling in the user info and vote status.
     *
     * @param comment The comment to be created. It contains the content of the comment, the id of the user who posted the comment, and the id of the entity to which the comment belongs.
     * @return The created comment with the user info and vote status filled in.
     */
    @Transactional
    public Comment createComment(Comment comment) {
        comment = commentRepo.save(comment);

        if (comment.getReplyTo() != null && comment.getReplyTo() != 0)
            commentRepo.increaseReplyCount(comment.getReplyTo());

        composeSingleComment(comment);
        return comment;
    }


    /**
     * Deletes a comment from the repository.
     * This annotation makes the method run within a transaction context.
     *
     * @param comment The comment to be deleted. It contains the id of the comment and the id of the user who posted the comment.
     */
    public void deleteComment(Comment comment) {
        commentRepo.deleteByIdAndUserId(comment.getId(), comment.getUserId());
    }


    /**
     * compose the comment, fill the user info and vote status
     *
     * @param comment the comment to be composed
     */
    private void composeSingleComment(Comment comment) {
        if (enableUser) {
            var token = Principal.currentToken();
            var user = userClient.getUser(comment.getUserId(), token);
            comment.setUser(user);
        }
    }

    /**
     * compose the comment, fill the user info and vote status
     *
     * @param comments the comments to be composed
     * @param userId   the id of the user who is requesting
     */
    private void composeComment(List<Comment> comments, Long userId) {
        if (comments.isEmpty()) {
            return;
        }
        var userMap = composeCommentAuthor(comments);
        var voted = composeCommentVoteStatus(comments, userId);
        flatComments(comments).forEach(comment -> {
            comment.setUser(userMap.get(comment.getUserId()));
            if (voted.contains(comment.getId())) {
                comment.setVoted(true);
            }
        });
    }


    /**
     * compose the comment, fill the user info
     *
     * @param comments the comments to be composed
     * @return the map of user id to user info
     */
    private Map<Long, User> composeCommentAuthor(List<Comment> comments) {
        if (!enableUser) {
            return Map.of();
        }
        // get user id list
        var userList = flatComments(comments).map(Comment::getUserId).distinct().toList();
        // fetch user info
        var token = Principal.currentToken();
        var users = userClient.getManyUser(userList, token);
        // covert to map
        return users.stream().collect(Collectors.toMap(User::getId, identity()));
    }

    /**
     * compose the comment, fill the vote status
     *
     * @param comments the comments to be composed
     * @param userId   the id of the user who is requesting
     * @return the set of comment id which the user has voted
     */
    private Set<Long> composeCommentVoteStatus(List<Comment> comments, Long userId) {
        if (!enableVote || userId == 0) {
            return Set.of();
        }
        var commentIdList = flatComments(comments).map(Comment::getId).toList();
        //  check if voted
        var votedIds = voteClient.queryInLikes(commentIdList, Principal.currentToken());
        // convert to set
        return Set.copyOf(votedIds);
    }

    /**
     * flat the comments and reply comments(only one level) to a stream
     *
     * @param comments the comments to be flatted
     * @return the stream of comments
     */
    private static Stream<Comment> flatComments(List<Comment> comments) {
        return comments.stream().flatMap(c -> {
            if (c.getReplyComments() == null) {
                return Stream.of(c);
            } else {
                return Stream.concat(Stream.of(c), c.getReplyComments().stream());
            }
        });
    }
}
