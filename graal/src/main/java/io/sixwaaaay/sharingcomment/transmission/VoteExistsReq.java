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
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import io.sixwaaaay.sharingcomment.util.ListLongRpcSerializer;
import io.sixwaaaay.sharingcomment.util.LongRpcSerializer;
import lombok.Data;

import java.util.List;

@Data
public class VoteExistsReq {
    private String type;
    @JsonProperty("subject_id")
    @JsonSerialize(using = LongRpcSerializer.class)
    private Long subjectId;
    @JsonProperty("target_type")
    private String targetType;
    @JsonProperty("target_ids")
    @JsonSerialize(using = ListLongRpcSerializer.class)
    private List<Long> targetIds;

    public VoteExistsReq(long subjectId, List<Long> targetIds) {
        this.type = "user";
        this.targetType = "comment";
        this.subjectId = subjectId;
        this.targetIds = targetIds;
    }
}
