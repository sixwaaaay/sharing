/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.util;

import io.sixwaaaay.sharingcomment.tools.JwtUtil;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
public class TokenTest {

    @Autowired
    private TokenParser tokenParser;

    @Autowired
    private JwtUtil jwtUtil;

    @Test
    public void testParseWithValidToken() {
        var l = 12314214521512131L;
        var id = String.valueOf(l);
        var johnDoe = "John Doe";
        var validToken = jwtUtil.generateToken(johnDoe, id);
        var principal = tokenParser.parse("Bearer " + validToken);
        assertTrue(principal.isPresent());
        assertEquals(johnDoe, principal.get().getName());
        assertEquals(l, principal.get().getId());
    }

    @Test
    public void testParseWithInvalidToken() {
        var invalidToken = "Invalid token";
        var principal = tokenParser.parse(invalidToken);
        assertFalse(principal.isPresent());
    }

    @Test
    public void testParseWithEmptyToken() {
        var emptyToken = "";
        var principal = tokenParser.parse(emptyToken);
        assertFalse(principal.isPresent());
    }
}