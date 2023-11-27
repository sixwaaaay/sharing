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
import io.sixwaaaay.sharingcomment.transmission.VoteExistsReq;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Limit;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class CommentService {
    private final CommentRepository commentRepo;

    private final VoteClient voteClient;

    private final UserClient userRepo;

    private final boolean enableVote;
    private final boolean enableUser;

    public CommentService(CommentRepository commentRepo, VoteClient voteClient, UserClient userRepo, @Value("${service.vote.enabled}") boolean enableVote, @Value("${service.user.enabled}") boolean enableUser) {
        this.commentRepo = commentRepo;
        this.voteClient = voteClient;
        this.userRepo = userRepo;
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
        var userList = comments.stream().flatMap(c -> c.getReplyComments().stream()).
                map(Comment::getUserId).distinct().toList();
        // fetch user info
        var users = userRepo.getManyUser(new GetMultipleUserReq(userList, userId));
        // covert to map
        var userMap = users.getUsers().stream().collect(Collectors.toMap(User::getId, user -> user));
        // fill user info
        comments.stream().flatMap(c -> c.getReplyComments().stream()).
                forEach(comment -> comment.setUser(userMap.get(comment.getUserId())));
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
        var commentIdList = comments.stream().flatMap(c -> c.getReplyComments().stream()).map(Comment::getId).toList();
        //  check if voted
        var voteExistsReply = voteClient.exists(new VoteExistsReq(userId, commentIdList));
        // convert to set
        var existedVote = voteExistsReply.getExists();

        // fill vote status
        comments.stream().flatMap(c -> c.getReplyComments().stream()).
                forEach(comment -> comment.setVoted(existedVote.contains(comment.getId())));
    }
}
