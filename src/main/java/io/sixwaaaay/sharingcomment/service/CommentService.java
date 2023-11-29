/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.service;


import io.sixwaaaay.sharingcomment.client.UserClient;
import io.sixwaaaay.sharingcomment.client.VoteClient;
import io.sixwaaaay.sharingcomment.domain.Comment;
import io.sixwaaaay.sharingcomment.domain.User;
import io.sixwaaaay.sharingcomment.repository.CommentRepository;
import io.sixwaaaay.sharingcomment.transmission.GetMultipleUserReq;
import io.sixwaaaay.sharingcomment.transmission.GetUserReq;
import io.sixwaaaay.sharingcomment.transmission.VoteExistsReq;
import io.sixwaaaay.sharingcomment.transmission.VoteReq;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Limit;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
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
     * get the main comment list
     *
     * @param belongTo the id of the object which the comments belong to
     * @param id       the id of the earliest comment in the previous page
     * @param userId   the id of the user who is requesting
     * @return page of comments
     */

    public List<Comment> getMainCommentList(Long belongTo, Long id, Integer size, Long userId) {
        var mainComments = commentRepo.findByBelongToAndIdLessThanAndReplyToNullOrderByIdDesc(belongTo, id, Limit.of(size));
//       for each main comment which has reply comments, get the latest 2 reply comments
        mainComments.stream().filter(comment -> comment.getReplyCount() != 0).forEach(comment -> {
            var replyComments = commentRepo.findByReplyToAndIdGreaterThanOrderByIdAsc(comment.getId(), 0L, Limit.of(2));
            comment.setReplyComments(replyComments);
        });

        composeComment(mainComments, userId);
        return mainComments;
    }

    /**
     * get the reply comment list
     *
     * @param replyTo the id of the comment which the comments belong to
     * @param id      the id of the latest comment in the previous page
     * @param userId  the id of the user who is requesting
     * @return the list of comments
     */
    public List<Comment> getReplyCommentList(Long replyTo, Long id, Integer size, Long userId) {
        var comments = commentRepo.findByReplyToAndIdGreaterThanOrderByIdAsc(replyTo, id, Limit.of(size));
        composeComment(comments, userId);
        return comments;
    }


    /**
     * create a comment
     *
     * @param comment the comment to be created
     * @return the created comment
     */
    @Transactional
    public Comment createComment(Comment comment) {
        comment = commentRepo.save(comment);
        if (comment.getReplyTo() != null && comment.getReplyTo() != 0) {
            commentRepo.increaseReplyCount(comment.getReplyTo());
        }
        composeSingleComment(comment, comment.getUserId());
        return comment;
    }


    /**
     * delete a comment
     *
     * @param comment the comment to be deleted
     */
    @Transactional
    public void deleteComment(Comment comment) {
        commentRepo.deleteByIdAndUserId(comment.getId(), comment.getUserId());
        if (comment.getReplyTo() != null && comment.getReplyTo() != 0) {
            commentRepo.decreaseReplyCount(comment.getReplyTo());
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
    }

    /**
     * cancel vote a comment
     *
     * @param userId    the id of the user who is requesting
     * @param commentId the id of the comment to be unvoted
     */
    public void cancelVoteComment(Long userId, Long commentId) {
        voteClient.itemDelete(new VoteReq(userId, commentId));
    }


    /**
     * compose the comment, fill the user info and vote status
     *
     * @param comment the comment to be composed
     * @param userId  the id of the user who is requesting
     */
    private void composeSingleComment(Comment comment, Long userId) {
        if (enableUser) {
            var user = userClient.getUser(new GetUserReq(comment.getUserId(), userId));
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
            composeCommentAuthor(comments, userId);
        if (enableVote)
            composeCommentVoteStatus(comments, userId);
    }


    /**
     * compose the comment, fill the user info
     *
     * @param comments the comments to be composed
     * @param userId   the id of the user who is requesting
     */
    private void composeCommentAuthor(List<Comment> comments, Long userId) {
        // get user id list
        var userList = flatComments(comments).map(Comment::getUserId).collect(Collectors.toUnmodifiableSet());
        // fetch user info
        var users = userClient.getManyUser(new GetMultipleUserReq(userList, userId));
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
