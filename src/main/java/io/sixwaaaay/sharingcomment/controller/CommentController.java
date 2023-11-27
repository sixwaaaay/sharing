/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.controller;


import io.sixwaaaay.sharingcomment.domain.Comment;
import io.sixwaaaay.sharingcomment.request.CommentRequest;
import io.sixwaaaay.sharingcomment.request.UserAuth;
import io.sixwaaaay.sharingcomment.request.error.NoUserExitsError;
import io.sixwaaaay.sharingcomment.service.CommentService;
import io.sixwaaaay.sharingcomment.util.TokenParser;
import jakarta.validation.Valid;
import org.springframework.web.bind.annotation.*;

import java.sql.Timestamp;
import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/comments")
public class CommentController {

    private final TokenParser tokenParser;

    private final CommentService commentService;


    public CommentController(TokenParser tokenParser, CommentService commentService) {
        this.tokenParser = tokenParser;
        this.commentService = commentService;
    }


    /**
     * main comment list also known as first level comment list
     *
     * @param belongTo the id of target object
     * @return the list of main comment
     */
    @GetMapping("/main")
    public List<Comment> getMainCommentList(
            @RequestParam("belongTo") Long belongTo,
            @RequestParam("page") Optional<Long> id,
            @RequestParam(value = "size", defaultValue = "10") Integer size,
            @RequestHeader(value = "Authorization", defaultValue = "") String header
    ) {
        var userAuth = tokenParser.parse(header);
        return commentService.getMainCommentList(belongTo, id.orElse(Long.MAX_VALUE), size, userAuth.map(UserAuth::getId).orElse(0L));
    }

    /**
     * reply comment list also known as second level comment list
     *
     * @param replyTo the id of target comment
     * @return the list of reply comment
     */
    @GetMapping("/reply")
    public List<Comment> getReplyCommentList(
            @RequestParam("replyTo") Long replyTo,
            @RequestParam("page") Optional<Long> id,
            @RequestParam(value = "size", defaultValue = "10") Integer size,
            @RequestHeader(value = "Authorization", defaultValue = "") String header
    ) {
        var userAuth = tokenParser.parse(header);
        return commentService.getReplyCommentList(replyTo, id.orElse(0L), size, userAuth.map(UserAuth::getId).orElse(0L));
    }

    /**
     * create a comment
     *
     * @return the created comment
     */
    @PostMapping
    public Comment createComment(@Valid @RequestBody CommentRequest request, @RequestHeader(value = "Authorization", defaultValue = "") String header) {
        var userAuth = tokenParser.parse(header);
        var comment = new Comment();
        comment.setUserId(userAuth.orElseThrow(NoUserExitsError::supply).getId()); // throw exception if userAuth is empty
        comment.setBelongTo(request.getBelongTo());
        comment.setContent(request.getContent());
        comment.setReplyTo(request.getReplyTo());
        comment.setCreatedAt(new Timestamp(System.currentTimeMillis()));

        comment = commentService.createComment(comment);
        return comment;
    }

    @DeleteMapping("/{id}")
    public void deleteComment(
            @PathVariable("id") Long id,
            @RequestHeader(value = "Authorization", defaultValue = "") String header,
            @RequestBody CommentRequest request
    ) {
        var userAuth = tokenParser.parse(header);
        var comment = new Comment();
        comment.setUserId(userAuth.orElseThrow(NoUserExitsError::supply).getId()); // throw exception if userAuth is empty
        comment.setId(id);
        comment.setReplyTo(request.getReplyTo());
        commentService.deleteComment(comment);
    }
}
