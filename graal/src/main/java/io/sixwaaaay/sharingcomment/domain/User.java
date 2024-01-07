/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.domain;


import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Builder;
import lombok.Data;

/**
 * The User class represents a user in the system.
 */
@Data
@Builder
public class User {
    /**
     * The unique ID of the user.
     */
    private Long id;

    /**
     * The name of the user.
     */
    private String name;

    /**
     * A flag indicating whether the user is followed by the current user.
     */
    @JsonProperty("is_follow")
    private Boolean isFollow;

    /**
     * The URL of the user's avatar image.
     */
    @JsonProperty("avatar_url")
    private String avatarUrl;

    /**
     * The URL of the user's background image.
     */
    @JsonProperty("bg_url")
    private String bgUrl;

    /**
     * The bio of the user.
     */
    private String bio;

    /**
     * The number of likes given by the user.
     */
    @JsonProperty("likes_given")
    private Integer likesGiven;

    /**
     * The number of likes received by the user.
     */
    @JsonProperty("likes_received")
    private Integer likesReceived;

    /**
     * The number of videos posted by the user.
     */
    @JsonProperty("videos_posted")
    private Integer videosPosted;

    /**
     * The number of users that the user is following.
     */
    private Integer following;

    /**
     * The number of followers of the user.
     */
    private Integer followers;
}