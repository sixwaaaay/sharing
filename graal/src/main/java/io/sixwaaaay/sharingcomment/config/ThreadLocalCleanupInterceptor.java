package io.sixwaaaay.sharingcomment.config;


import org.springframework.web.servlet.HandlerInterceptor;

import io.sixwaaaay.sharingcomment.util.DbContext;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;


public class ThreadLocalCleanupInterceptor implements HandlerInterceptor {

    /**
     * execute clean up operation after request.
     * remove dbcontext threadlocal value to prevent memory leak.
     */
    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex)
            throws Exception {
        DbContext.clear();
    }
}