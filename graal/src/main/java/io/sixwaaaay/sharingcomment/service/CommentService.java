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
import io.sixwaaaay.sharingcomment.repository.CountRepository;
import io.sixwaaaay.sharingcomment.request.Principal;
import io.sixwaaaay.sharingcomment.transmission.VoteExistsReq;
import io.sixwaaaay.sharingcomment.transmission.VoteReq;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Limit;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import static io.sixwaaaay.sharingcomment.util.TokenParser.principal;
import static java.util.function.Function.identity;

@Service
public class CommentService {
    private final CommentRepository commentRepo;

    private final CountRepository countRepo;

    private final VoteClient voteClient;

    private final UserClient userClient;

    private final boolean enableVote;
    private final boolean enableUser;

    public CommentService(CommentRepository commentRepo, CountRepository countRepo, VoteClient voteClient, UserClient userClient, @Value("${service.vote.enabled}") boolean enableVote, @Value("${service.user.enabled}") boolean enableUser) {
        this.commentRepo = commentRepo;
        this.countRepo = countRepo;
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
        var result = new CommentResult();
        var count = countRepo.findById(belongTo).orElse(new Count(belongTo, 0)).getCommentCount();

        result.setAllCount(count);
        if (count == 0) {
            result.setComments(List.of());
            return result;
        }

        var mainComments = commentRepo.findByBelongToAndIdLessThanAndReplyToNullOrderByIdDesc(belongTo, id, Limit.of(size));
//       for each main comment which has reply comments, get the latest 2 reply comments
        mainComments.stream().filter(comment -> comment.getReplyCount() != 0).forEach(comment -> {
            var replyComments = commentRepo.findByBelongToAndReplyToAndIdGreaterThanOrderByIdAsc(belongTo, comment.getId(), 0L, Limit.of(2));
            comment.setReplyComments(replyComments);
        });

        composeComment(mainComments, userId);
        result.setComments(mainComments);
        if (mainComments.size() == size) {
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
        var comments = commentRepo.findByBelongToAndReplyToAndIdGreaterThanOrderByIdAsc(belongTo, replyTo, id, Limit.of(size));
        composeComment(comments, userId);
        var result = new ReplyResult();
        result.setComments(comments);
        if (comments.size() == size) {
            result.setNextPage(comments.getLast().getId());
        }
        return result;
    }

    /**
     * This method is used to create a new comment in the system.
     * It first saves the comment to the repository.
     * Then, it increases the count of comments belonging to the same entity.
     * If the count does not exist, it creates a new count with a value of 1.
     * If the new comment is a reply to another comment, it increases the reply count of the original comment.
     * Finally, it composes the comment by filling in the user info and vote status.
     *
     * @param comment The comment to be created. It contains the content of the comment, the id of the user who posted the comment, and the id of the entity to which the comment belongs.
     * @return The created comment with the user info and vote status filled in.
     */
    @Transactional
    public Comment createComment(Comment comment) {
        comment = commentRepo.save(comment);
        var updated = countRepo.increaseCount(comment.getBelongTo());
        if (!updated)
            countRepo.createCount(comment.getBelongTo()); // create count if not exist
        if (comment.getReplyTo() != null && comment.getReplyTo() != 0)
            commentRepo.increaseReplyCount(comment.getReplyTo());
        composeSingleComment(comment);
        return comment;
    }


    /**
     * Deletes a comment from the repository.
     * If the comment is successfully deleted, it decreases the count of comments belonging to the same entity.
     * If the deleted comment is a reply to another comment, it decreases the reply count of the original comment.
     * This annotation makes the method run within a transaction context.
     *
     * @param comment The comment to be deleted. It contains the id of the comment and the id of the user who posted the comment.
     */
    @Transactional
    public void deleteComment(Comment comment) {
        var deleted = commentRepo.deleteByIdAndUserId(comment.getId(), comment.getUserId());
        if (deleted) {
            countRepo.decreaseCount(comment.getBelongTo());
            if (comment.getReplyTo() != null && comment.getReplyTo() != 0) {
                commentRepo.decreaseReplyCount(comment.getReplyTo());
            }
        }
    }

    /**
     * vote a comment
     *
     * @param userId    the id of the user who is requesting
     * @param commentId the id of the comment to be voted
     */
    public void voteComment(long userId, long commentId) {
        voteClient.itemAdd(new VoteReq(userId, commentId));
        commentRepo.increaseLikeCount(commentId);
    }

    /**
     * cancel vote a comment
     *
     * @param userId    the id of the user who is requesting
     * @param commentId the id of the comment to be unvoted
     */
    public void cancelVoteComment(Long userId, Long commentId) {
        voteClient.itemDelete(new VoteReq(userId, commentId));
        commentRepo.decreaseLikeCount(commentId);
    }


    /**
     * compose the comment, fill the user info and vote status
     *
     * @param comment the comment to be composed
     */
    private void composeSingleComment(Comment comment) {
        if (enableUser) {
            var token = principal.get().map(Principal::getToken).orElse("");
            var user = userClient.getUser(comment.getUserId(), token);
            comment.setUser(user.getUser());
        }
    }

    /**
     * compose the comment, fill the user info and vote status
     *
     * @param comments the comments to be composed
     * @param userId   the id of the user who is requesting
     */
    private void composeComment(List<Comment> comments, Long userId) {
        if (enableUser)
            composeCommentAuthor(comments);
        if (enableVote)
            composeCommentVoteStatus(comments, userId);
    }


    /**
     * compose the comment, fill the user info
     *
     * @param comments the comments to be composed
     */
    private void composeCommentAuthor(List<Comment> comments) {
        // get user id list
        var userList = flatComments(comments).map(Comment::getUserId).distinct().toList();
        if (userList.isEmpty()) {
            return;
        } // return if no user id, avoid unnecessary request and missing parameter error
        // fetch user info
        var token = principal.get().map(Principal::getToken).orElse("");
        var users = userClient.getManyUser(userList, token);
        // covert to map
        var userMap = users.getUsers().stream().collect(Collectors.toMap(User::getId, identity()));
        // fill user info
        flatComments(comments).forEach(comment -> comment.setUser(userMap.get(comment.getUserId())));
    }

    /**
     * compose the comment, fill the vote status
     *
     * @param comments the comments to be composed
     * @param userId   the id of the user who is requesting
     */
    private void composeCommentVoteStatus(List<Comment> comments, Long userId) {
        if (userId == 0) {
            return;
        }
        var commentIdList = flatComments(comments).map(Comment::getId).toList();
        if (commentIdList.isEmpty()) {
            return;
        } // return if no comment id, avoid unnecessary request and missing parameter error
        //  check if voted
        var voteExistsReply = voteClient.exists(new VoteExistsReq(userId, commentIdList));
        // convert to set
        var existedVote = voteExistsReply.getExists();
        // fill vote status
        flatComments(comments).forEach(comment -> comment.setVoted(existedVote.contains(comment.getId())));
    }

    /**
     * flat the comments and reply comments(only one level) to a stream
     *
     * @param comments the comments to be flatted
     * @return the stream of comments
     */
    private static Stream<Comment> flatComments(List<Comment> comments) {
        return comments.stream().flatMap(c -> {
            if (c.getReplyComments() == null)
                return Stream.of(c);
            else
                return Stream.concat(Stream.of(c), c.getReplyComments().stream());
        });
    }
}
