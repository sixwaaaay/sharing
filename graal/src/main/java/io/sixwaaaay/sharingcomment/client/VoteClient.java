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

package io.sixwaaaay.sharingcomment.client;


import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.service.annotation.PostExchange;

import java.util.List;


public interface VoteClient {

    /**
     * query whether the objectIds(aka commentIds) is liked.
     * @param objectIds the id list of the object(aka comment)
     * @param token the token of the user
     * @return the list of the objectIds which is liked
     */
    @PostExchange("/graph/comments/likes")
    List<Long> queryInLikes(@RequestBody List<Long> objectIds, @RequestHeader(value = "Authorization", required = false) String token);
}
