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
