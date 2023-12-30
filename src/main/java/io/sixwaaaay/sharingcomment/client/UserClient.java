/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.client;


import io.sixwaaaay.sharingcomment.transmission.GetMultipleUserReply;
import io.sixwaaaay.sharingcomment.transmission.GetUserReply;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.service.annotation.GetExchange;

import java.util.Collection;


public interface UserClient {
    @GetExchange("/users/{user_id}")
    GetUserReply getUser(@PathVariable("user_id") long id, @RequestHeader(value = "Authorization", required = false) String token);

    @GetExchange("/users")
    GetMultipleUserReply getManyUser(@RequestParam("ids") Collection<Long> ids, @RequestHeader(value = "Authorization", required = false) String token);
}