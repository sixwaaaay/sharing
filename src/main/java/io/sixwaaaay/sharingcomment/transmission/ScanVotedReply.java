/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import lombok.Data;

import java.util.List;

@Data
public class ScanVotedReply {
    private List<Long> targetIds;
}
