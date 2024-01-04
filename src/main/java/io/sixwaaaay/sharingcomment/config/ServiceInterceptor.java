/*
 * Copyright (c) 2024 sixwaaaay
 *  All rights reserved.
 */

package io.sixwaaaay.sharingcomment.config;

import io.sixwaaaay.sharingcomment.util.TokenParser;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;


@Component
@AllArgsConstructor
public class ServiceInterceptor implements HandlerInterceptor {

    private final TokenParser tokenParser;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) {
        var token = request.getHeader("Authorization");
        tokenParser.parse(token);
        return true;
    }
}
