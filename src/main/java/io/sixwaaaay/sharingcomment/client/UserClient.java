/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.client;


import io.sixwaaaay.sharingcomment.transmission.GetUserReply;
import io.sixwaaaay.sharingcomment.transmission.GetUserReq;
import io.sixwaaaay.sharingcomment.transmission.GetMultipleUserReply;
import io.sixwaaaay.sharingcomment.transmission.GetMultipleUserReq;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.service.annotation.PostExchange;



public interface UserClient {
    @PostExchange("/sixwaaaay.user.UserService/GetUser")
    GetUserReply getUser(@RequestBody GetUserReq req);

    @PostExchange("/sixwaaaay.user.UserService/GetUsers")
    GetMultipleUserReply getManyUser(@RequestBody GetMultipleUserReq req);
}
