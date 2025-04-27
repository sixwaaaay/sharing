package io.sixwaaaay.sharingcomment.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class WebConfig implements WebMvcConfigurer {

    /**
     * add interceptor to interceptor registry.
     * configure global thread local variable cleanup interceptor to handle the
     * cleanup work of thread local variables.
     */
    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(new ThreadLocalCleanupInterceptor())
                .addPathPatterns("/**"); // include all requests
    }
}