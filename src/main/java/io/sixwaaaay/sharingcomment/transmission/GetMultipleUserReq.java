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
public class GetMultipleUserReq {
    @JsonProperty("user_ids")
    private List<Long> userIds;
    @JsonProperty("subject_id")
    private Long subjectId;
}
