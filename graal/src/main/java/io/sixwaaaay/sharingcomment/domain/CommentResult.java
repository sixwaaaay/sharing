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

package io.sixwaaaay.sharingcomment.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

/**
 * This class represents the result of a comment query.
 * It includes the total count of comments, the previous and next page identifiers, and a list of comments.
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class CommentResult {

    /**
     * The total count of all comments.
     */
    @JsonProperty("all_count")
    private int allCount;

    /**
     * The identifier of the next page. Null if there is no next page.
     */
    @JsonProperty("next_page")
    private Long nextPage;

    /**
     * A list of comments for the current page.
     */
    @JsonProperty("comments")
    private List<Comment> comments;
}