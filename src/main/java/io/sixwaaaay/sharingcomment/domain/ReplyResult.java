/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.domain;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class ReplyResult {
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
