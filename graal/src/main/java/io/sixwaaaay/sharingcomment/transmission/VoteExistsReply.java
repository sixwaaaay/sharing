/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import lombok.Data;

import java.util.Set;

@Data
public class VoteExistsReply {
    private Set<Long> exists;
}
