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

    public static ThreadLocal<Optional<Principal>> principal = new ThreadLocal<>();

    public TokenParser(@Value("${jwt.secret}") String secret) {



        var secretKey = Keys.hmacShaKeyFor(secret.getBytes());
        jwtParser = Jwts.parser()
                .verifyWith(secretKey)
                .build();
    }

    /**
     * Parse a Bearer token and return a Principal object
     *
     * @param token Bearer token
     * @return Optional of Principal
     */
    public Optional<Principal> parse(String token) {
        var PREFIX = "Bearer ";
        if (token == null || token.isEmpty() || !token.startsWith(PREFIX)) {
            principal.set(Optional.empty());
            return Optional.empty();
        }

        var tokenString = token.substring(PREFIX.length());
        var claimsJws = jwtParser.parseSignedClaims(tokenString);
        var name = claimsJws.getPayload().get("name", String.class);
        var id = claimsJws.getPayload().get("id", String.class);
        var value = new Principal(name, Long.parseLong(id), token);
        principal.set(Optional.of(value));
        return Optional.of(value);
    }

}
