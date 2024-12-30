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

package io.sixwaaaay.sharingcomment.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.annotation.ReadOnlyProperty;
import org.springframework.data.annotation.Transient;
import org.springframework.data.relational.core.mapping.Column;
import org.springframework.data.relational.core.mapping.Table;

import java.io.Serial;
import java.io.Serializable;
import java.time.LocalDateTime;
import java.util.List;

/**
 * This class represents a Comment in the system.
 * It includes various properties related to a comment such as id, user id, content, reply to, belong to, created at, reply count, like count, reply comments, voted and user.
 * This class implements Serializable interface for object serialization.
 */
@Table("comments")
@Data
@NoArgsConstructor
public class Comment implements Serializable {

    /**
     * The unique identifier for the comment.
     */
    @Serial
    private static final long serialVersionUID = 1L;

    /**
     * The unique identifier for the comment.
     */
    @Id
    private long id;

    /**
     * The unique identifier of the user who posted the comment.
     */
    @Column("user_id")
    @JsonProperty("user_id")
    private long userId;

    /**
     * The content of the comment.
     */
    private String content;

    /**
     * The unique identifier of the comment to which this comment is a reply. Null if this comment is not a reply.
     */
    @Column("reply_to")
    @JsonProperty("reply_to")
    private Long replyTo;

    /**
     * The unique identifier of the parent comment.
     */
    @Column("refer_to")
    @JsonProperty("refer_to")
    private Long referTo;

    /**
     * The unique identifier of the entity to which this comment belongs.
     */
    @Column("belong_to")
    @JsonProperty("belong_to")
    private long belongTo;

    /**
     * The timestamp when the comment was created.
     */
    @Column("created_at")
    @JsonProperty("created_at")
    @ReadOnlyProperty
    private LocalDateTime createdAt;

    /**
     * The count of replies to this comment.
     */
    @Column("reply_count")
    @JsonProperty("reply_count")
    @ReadOnlyProperty
    private int replyCount;

    /**
     * The count of likes for this comment.
     */
    @Column("like_count")
    @JsonProperty("like_count")
    @ReadOnlyProperty
    private int likeCount;

    /**
     * A list of comments which are replies to this comment. This field is not persisted in the database.
     */
    @Transient
    @JsonProperty("reply_comments")
    private List<Comment> replyComments;

    /**
     * A boolean indicating whether the current user has voted for this comment. This field is not persisted in the database.
     */
    @Transient
    private boolean voted;

    /**
     * The user who posted this comment. This field is not persisted in the database.
     */
    @Transient
    private User user;
}

