/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.util;

import io.jsonwebtoken.JwtParser;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;
import io.sixwaaaay.sharingcomment.request.Principal;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

import java.util.Optional;


@Component
public class TokenParser {
    private final JwtParser jwtParser;

    public TokenParser(@Value("${jwt.secret}") String secret) {
        var secretKey = Keys.hmacShaKeyFor(secret.getBytes());
        jwtParser = Jwts.parser()
                .verifyWith(secretKey)
                .build();
    }

    /**
     * Parse a token and return a Principal object
     * @param token Bearer token
     * @return Optional of Principal
     */
    public Optional<Principal> parse(String token) {
        if (token == null || token.isEmpty()) return Optional.empty();

        var claimsJws = jwtParser.parseSignedClaims(token);
        var name = claimsJws.getPayload().get("name", String.class);
        var id = claimsJws.getPayload().get("id", String.class);
        return Optional.of(new Principal(name, Long.parseLong(id)));
    }

}
