/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import lombok.Data;

@Data
public class ScanVotedReq {
    private Integer limit;
    private Long subjectId;
    private String targetType;
    private Long token;
    private String type;
}
