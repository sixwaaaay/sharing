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


import io.sixwaaaay.sharingcomment.domain.User;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.service.annotation.GetExchange;

import java.util.Collection;
import java.util.List;


public interface UserClient {
    @GetExchange("/users/{user_id}")
    User getUser(@PathVariable("user_id") long id, @RequestHeader(value = "Authorization", required = false) String token);

    @GetExchange("/users")
    List<User> getManyUser(@RequestParam("ids") Collection<Long> ids, @RequestHeader(value = "Authorization", required = false) String token);
}