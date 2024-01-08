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

package io.sixwaaaay.sharingcomment.controller;


import io.sixwaaaay.sharingcomment.domain.Comment;
import io.sixwaaaay.sharingcomment.domain.CommentResult;
import io.sixwaaaay.sharingcomment.domain.ReplyResult;
import io.sixwaaaay.sharingcomment.request.CommentRequest;
import io.sixwaaaay.sharingcomment.request.Principal;
import io.sixwaaaay.sharingcomment.request.error.NoUserExitsError;
import io.sixwaaaay.sharingcomment.service.CommentService;
import jakarta.validation.Valid;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.Optional;

import static io.sixwaaaay.sharingcomment.util.TokenParser.principal;

@RestController
@RequestMapping("/comments")
@AllArgsConstructor
public class CommentController {

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
            @RequestParam(value = "page") Optional<Long> id,
            @RequestParam(value = "size", defaultValue = "10") Integer size
    ) {
        var userId = principal.get().map(Principal::getId).orElse(0L);
        return commentService.getMainCommentList(belongTo, id.orElse(Long.MAX_VALUE), size, userId);
    }

    /**
     * reply comment list also known as second level comment list
     *
     * @param replyTo the id of target comment
     * @return the list of reply comment
     */
    @GetMapping("/reply")
    public ReplyResult getReplyCommentList(
            @RequestParam("belong_to") Long belongTo,
            @RequestParam("reply_to") Long replyTo,
            @RequestParam("page") Optional<Long> id,
            @RequestParam(value = "size", defaultValue = "10") Integer size
    ) {
        var userId = principal.get().map(Principal::getId).orElse(0L);
        return commentService.getReplyCommentList(belongTo, replyTo, id.orElse(0L), size, userId);
    }

    /**
     * create a comment
     *
     * @return the created comment
     */
    @PostMapping
    public Comment createComment(@Valid @RequestBody CommentRequest request) {
        var comment = new Comment();
        var id = principal.get().map(Principal::getId).orElseThrow(NoUserExitsError::supply);
        comment.setUserId(id); // throw exception if principal is empty
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
            @RequestBody CommentRequest request
    ) {
        var userId = principal.get().map(Principal::getId).orElseThrow(NoUserExitsError::supply);
        var comment = new Comment();
        comment.setUserId(userId); // throw exception if principal is empty
        comment.setId(id);
        comment.setReplyTo(request.getReplyTo());
        commentService.deleteComment(comment);
    }

    /**
     * vote a comment
     *
     * @param id the id of comment
     */
    @PostMapping("/action/like/{id}")
    public void voteComment(
            @PathVariable long id) {
        var userId = principal.get().map(Principal::getId).orElseThrow(NoUserExitsError::supply);
        commentService.voteComment(userId, id);
    }

    /**
     * cancel vote a comment
     *
     * @param id the id of comment
     */
    @DeleteMapping("/action/like/{id}")
    public void cancelVoteComment(
            @PathVariable long id) {
        var userId = principal.get().map(Principal::getId).orElseThrow(NoUserExitsError::supply);
        commentService.cancelVoteComment(userId, id);
    }
}
