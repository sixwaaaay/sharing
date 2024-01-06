/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import io.sixwaaaay.sharingcomment.domain.User;
import lombok.Data;

import java.util.List;

@Data
public class GetMultipleUserReply {
    private List<User> users;
}
