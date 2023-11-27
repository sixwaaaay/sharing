/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Data;
import org.hibernate.validator.constraints.Range;

@Data

public class VoteReq {
    private String type;
    @NotNull
    @Range(min = 1, message = "subject_id must be greater than 0")
    @JsonProperty("subject_id")
    private Long subjectId;
    @JsonProperty("target_type")
    private String targetType;
    @NotNull
    @Range(min = 1, message = "target_id must be greater than 0")
    @JsonProperty("target_id")
    private Long targetId;

    public VoteReq(long subjectId, long targetId) {
        this.type = "user";
        this.targetType = "comment";
        this.subjectId = subjectId;
        this.targetId = targetId;
    }
}
