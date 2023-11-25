/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.client;

import io.sixwaaaay.sharingcomment.transmission.GetUserReq;
import io.sixwaaaay.sharingcomment.transmission.GetMultipleUserReq;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.List;

@SpringBootTest
public class UserClientTests {
    @Autowired
    private UserClient userClient;

    @Test
    void testGetUser() {
        var userReply = userClient.getUser(new GetUserReq(1L, 1L));
        Assertions.assertNotNull(userReply);
    }

    @Test
    void testGetUsers() {
        var usersReply = userClient.getManyUser(new GetMultipleUserReq(List.of(457232417502052951L, 457121784278309633L), 1L));
        Assertions.assertNotNull(usersReply);
    }
}
