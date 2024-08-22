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

package io.sixwaaaay.sharingcomment.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.Valid;
import jakarta.validation.constraints.AssertTrue;
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
     * The type of the comment.
     */
    @NotNull
    @JsonProperty("type")
    private String type;

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
     * The ID of the comment that this comment refers to.
     */
    @JsonProperty("refer_to")
    private Long referTo;

    /**
     * The ID of the entity to which the comment belongs.
     * It must be greater than 0.
     */
    @NotNull(message = "belong_to must not be null")
    @Range(min = 1, message = "belong_to must be greater than 0")
    @JsonProperty("belong_to")
    private Long belongTo;

    /**
     * reply_to and parent_id must be both null or not null at the same time
     */
    @AssertTrue(message = "reply_to and refer_to must be both null or not null at the same time")
    public boolean isValid() {
        return (replyTo == null && referTo == null) || (replyTo != null && referTo != null);
    }

    @AssertTrue(message = "type must be one of 'chore', 'default', 'video', 'post', 'music'")
    public boolean isValidType() {
        return type.equals("chore") || type.equals("default") || type.equals("video") || type.equals("post") || type.equals("music");
    }
}
