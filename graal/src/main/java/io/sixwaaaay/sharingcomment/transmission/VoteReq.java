/*
 * Copyright (c) 2023-2024 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.sixwaaaay.sharingcomment.transmission;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.NotNull;
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
