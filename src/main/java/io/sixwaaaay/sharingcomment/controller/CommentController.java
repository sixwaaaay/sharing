/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.controller;


import io.sixwaaaay.sharingcomment.domain.Comment;
import io.sixwaaaay.sharingcomment.domain.CommentResult;
import io.sixwaaaay.sharingcomment.domain.ReplyResult;
import io.sixwaaaay.sharingcomment.request.CommentRequest;
import io.sixwaaaay.sharingcomment.request.Principal;
import io.sixwaaaay.sharingcomment.request.error.NoUserExitsError;
import io.sixwaaaay.sharingcomment.service.CommentService;
import io.sixwaaaay.sharingcomment.util.TokenParser;
import jakarta.validation.Valid;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.Optional;

@RestController
@RequestMapping("/comments")
@AllArgsConstructor
public class CommentController {

    private final TokenParser tokenParser;

    private final CommentService commentService;

    /**
     * main comment list also known as first level comment list
     *
     * @param belongTo the id of target object
     * @return the list of main comment
     */
    @GetMapping("/main")
    public CommentResult getMainCommentList(
            @RequestParam("belong_to") Long belongTo,
            @RequestParam("page") Optional<Long> id,
            @RequestParam(value = "size", defaultValue = "10") Integer size,
            @RequestHeader(value = "Authorization", required = false) String header
    ) {
        var principal = tokenParser.parse(header);
        return commentService.getMainCommentList(belongTo, id.orElse(Long.MAX_VALUE), size, principal.map(Principal::getId).orElse(0L));
    }

    /**
     * reply comment list also known as second level comment list
     *
     * @param replyTo the id of target comment
     * @return the list of reply comment
     */
    @GetMapping("/reply")
    public ReplyResult getReplyCommentList(
            @RequestParam("reply_to") Long replyTo,
            @RequestParam("page") Optional<Long> id,
            @RequestParam(value = "size", defaultValue = "10") Integer size,
            @RequestHeader(value = "Authorization", required = false) String header
    ) {
        var principal = tokenParser.parse(header);
        return commentService.getReplyCommentList(replyTo, id.orElse(0L), size, principal.map(Principal::getId).orElse(0L));
    }

    /**
     * create a comment
     *
     * @return the created comment
     */
    @PostMapping
    public Comment createComment(@Valid @RequestBody CommentRequest request, @RequestHeader(value = "Authorization") String header) {
        var principal = tokenParser.parse(header);
        var comment = new Comment();
        comment.setUserId(principal.orElseThrow(NoUserExitsError::supply).getId()); // throw exception if principal is empty
        comment.setBelongTo(request.getBelongTo());
        comment.setContent(request.getContent());
        comment.setReplyTo(request.getReplyTo());
        var epochSecond = System.currentTimeMillis() / 1000;
        comment.setCreatedAt(LocalDateTime.ofEpochSecond(epochSecond, 0, ZoneOffset.ofHours(8)));

        comment = commentService.createComment(comment);
        return comment;
    }

    @DeleteMapping("/{id}")
    public void deleteComment(
            @PathVariable("id") Long id,
            @RequestHeader(value = "Authorization") String header,
            @RequestBody CommentRequest request
    ) {
        var principal = tokenParser.parse(header);
        var comment = new Comment();
        comment.setUserId(principal.orElseThrow(NoUserExitsError::supply).getId()); // throw exception if principal is empty
        comment.setId(id);
        comment.setReplyTo(request.getReplyTo());
        commentService.deleteComment(comment);
    }

    /**
     * vote a comment
     *
     * @param id     the id of comment
     * @param header the header of request
     */
    @PostMapping("/action/like/{id}")
    public void voteComment(
            @PathVariable long id,
            @RequestHeader(value = "Authorization") String header) {
        var principal = tokenParser.parse(header);
        commentService.voteComment(principal.map(Principal::getId).orElseThrow(NoUserExitsError::supply), id);
    }

    /**
     * cancel vote a comment
     *
     * @param id     the id of comment
     * @param header the header of request
     */
    @DeleteMapping("/action/like/{id}")
    public void cancelVoteComment(
            @PathVariable long id,
            @RequestHeader(value = "Authorization") String header) {
        var principal = tokenParser.parse(header);
        commentService.cancelVoteComment(principal.map(Principal::getId).orElseThrow(NoUserExitsError::supply), id);
    }
}
