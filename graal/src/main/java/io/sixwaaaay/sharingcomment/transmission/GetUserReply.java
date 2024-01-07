/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.annotation.JsonProperty;
import io.sixwaaaay.sharingcomment.domain.User;
import lombok.Data;

@Data
public class GetUserReply {
    @JsonProperty("user")
    private User user;
}
