/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
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

import java.sql.Timestamp;
import java.util.List;

@Table("comments")
@Data
@NoArgsConstructor
public class Comment {
    @Id
    private Long id;
    @Column("user_id")
    @JsonProperty("user_id")
    private Long userId;
    private String content;

    @Column("reply_to")
    @JsonProperty("reply_to")
    private Long replyTo;
    @Column("belong_to")
    @JsonProperty("belong_to")
    private Long belongTo;

    @Column("created_at")
    @JsonProperty("created_at")
    @ReadOnlyProperty
    private Timestamp createdAt;

    @Column("reply_count")
    @JsonProperty("reply_count")
    @ReadOnlyProperty
    private Integer replyCount;

    @Column("like_count")
    @JsonProperty("like_count")
    @ReadOnlyProperty
    private Integer likeCount;


    // fields which are not in the table
    @Transient
    @JsonProperty("reply_comments")
    private List<Comment> replyComments;
    @Transient
    private Boolean voted;
    @Transient
    private User user;
}

