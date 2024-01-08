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


import io.sixwaaaay.sharingcomment.transmission.*;
import jakarta.validation.Valid;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.service.annotation.PostExchange;


public interface VoteClient {
    @PostExchange("/item/add")
    VoteReply itemAdd(@Valid @RequestBody VoteReq req);

    @PostExchange("/item/delete")
    VoteReply itemDelete(@Valid @RequestBody VoteReq req);

    @PostExchange("/item/exists")
    VoteExistsReply exists(@Valid @RequestBody VoteExistsReq req);

    @PostExchange("/item/scan")
    ScanVotedReply scan(@RequestBody ScanVotedReq req);
}
