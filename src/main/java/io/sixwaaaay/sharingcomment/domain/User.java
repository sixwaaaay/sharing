/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.domain;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class User {
    private Long id;
    private String name;
    @JsonProperty("is_follow")
    private Boolean isFollow;
    @JsonProperty("avatar_url")
    private String avatarUrl;
    @JsonProperty("bg_url")
    private String bgUrl;
    private String bio;
    @JsonProperty("likes_given")
    private Integer likesGiven;
    @JsonProperty("likes_received")
    private Integer likesReceived;
    @JsonProperty("videos_posted")
    private Integer videosPosted;
    private Integer following;
    private Integer followers;
}
