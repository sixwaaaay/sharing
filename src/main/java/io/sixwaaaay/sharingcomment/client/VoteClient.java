/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.client;


import io.sixwaaaay.sharingcomment.transmission.*;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.service.annotation.PostExchange;


public interface VoteClient {
    @PostExchange("/item/add")
    VoteReply itemAdd(@RequestBody VoteReq req);

    @PostExchange("/item/delete")
    VoteReply itemDelete(@RequestBody VoteReq req);

    @PostExchange("/item/exists")
    VoteExistsReply exists(@RequestBody VoteExistsReq req);

    @PostExchange("/item/scan")
    ScanVotedReply scan(@RequestBody ScanVotedReq req);
}
