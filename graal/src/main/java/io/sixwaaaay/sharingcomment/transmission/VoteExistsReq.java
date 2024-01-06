/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import java.util.List;

@Data
public class VoteExistsReq {
    private String type;
    @JsonProperty("subject_id")
    private Long subjectId;
    @JsonProperty("target_type")
    private String targetType;
    @JsonProperty("target_ids")
    private List<Long> targetIds;

    public VoteExistsReq(long subjectId, List<Long> targetIds) {
        this.type = "user";
        this.targetType = "comment";
        this.subjectId = subjectId;
        this.targetIds = targetIds;
    }
}
