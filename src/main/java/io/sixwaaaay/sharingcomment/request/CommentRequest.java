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

@Data
@Valid
public class CommentRequest {
    @NotNull
    @Length(min = 1, max = 1000, message = "content length must be between 1 and 1000")
    private String content;
    @JsonProperty("reply_to")
    private Long replyTo;
    @NotNull
    @Range(min = 1, message = "belong_to must be greater than 0")
    @JsonProperty("belong_to")
    private Long belongTo;
}


