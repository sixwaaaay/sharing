/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class VoteReq {
    private String type;
    @JsonProperty("subject_id")
    private Long subjectId;
    @JsonProperty("target_type")
    private String targetType;
    @JsonProperty("target_id")
    private Long targetId;
}
