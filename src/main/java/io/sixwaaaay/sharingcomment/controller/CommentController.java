/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.controller;


import io.sixwaaaay.sharingcomment.domain.Comment;
import io.sixwaaaay.sharingcomment.request.UserAuth;
import io.sixwaaaay.sharingcomment.service.CommentService;
import io.sixwaaaay.sharingcomment.util.TokenParser;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/comments")
@Slf4j
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

}
