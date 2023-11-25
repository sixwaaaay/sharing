/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.List;

@Data
@AllArgsConstructor
public class VoteExistsReq {
    private String type;
    @JsonProperty("subject_id")
    private Long subjectId;
    @JsonProperty("target_type")
    private String targetType;
    @JsonProperty("target_ids")
    private List<Long> targetIds;
}
