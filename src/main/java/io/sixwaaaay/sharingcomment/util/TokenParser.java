/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.util;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.sixwaaaay.sharingcomment.request.Principal;
import org.springframework.stereotype.Component;

import java.util.Base64;
import java.util.Optional;


@Component
public class TokenParser {

    private final ObjectMapper objectMapper = new ObjectMapper();

    public TokenParser() {
        objectMapper.configure(com.fasterxml.jackson.databind.DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
    }

    public Optional<Principal> parse(String token) {
        try {
            byte[] authorizations = Optional.ofNullable(token)
                    .map(s -> s.split("\\."))
                    .filter(strings -> strings.length == 3)
                    .map(strings -> strings[1])
                    .map(Base64.getDecoder()::decode).orElseGet(() -> new byte[0]);
            return Optional.ofNullable(objectMapper.readValue(authorizations, Principal.class));
        } catch (Exception e) {
            return Optional.empty();
        }
    }

}
