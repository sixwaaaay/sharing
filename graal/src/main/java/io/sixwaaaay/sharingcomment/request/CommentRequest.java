/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.Valid;
import jakarta.validation.constraints.NotNull;
import lombok.Data;
import org.hibernate.validator.constraints.Length;
import org.hibernate.validator.constraints.Range;

/**
 * The CommentRequest class represents a request to create a new comment.
 */
@Data
@Valid
public class CommentRequest {
    /**
     * The content of the comment.
     * It must be between 1 and 1000 characters long.
     */
    @NotNull
    @Length(min = 1, max = 1000, message = "content length must be between 1 and 1000")
    private String content;

    /**
     * The ID of the comment to which this comment is a reply.
     * It can be null if the comment is not a reply.
     */
    @JsonProperty("reply_to")
    private Long replyTo;

    /**
     * The ID of the entity to which the comment belongs.
     * It must be greater than 0.
     */
    @NotNull(message = "belong_to must not be null")
    @Range(min = 1, message = "belong_to must be greater than 0")
    @JsonProperty("belong_to")
    private Long belongTo;
}
