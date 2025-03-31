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

package io.sixwaaaay.sharingcomment.config;

import io.sixwaaaay.sharingcomment.util.TokenParser;
import lombok.SneakyThrows;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpMethod;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configurers.AbstractHttpConfigurer;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

@Configuration
public class SecurityConfig {


    /**
     * configure user detail service
     *
     * @return user detail service
     */
    @Bean
    UserDetailsService emptyDetailsService() {
        return username -> {
            throw new UsernameNotFoundException("no local users, only JWT tokens allowed");
        };
    }


    /**
     * configure security filter chain
     *
     * @param http        http security builder
     * @param tokenParser token parser
     * @return security filter chain
     */
    @Bean
    @SneakyThrows
    public SecurityFilterChain filterChain(HttpSecurity http, TokenParser tokenParser) {
        http.authorizeHttpRequests((authorize) -> authorize
                        .requestMatchers(
                                HttpMethod.POST,
                                "/comments",
                                "/comments/action/like/*"
                        ).hasAuthority("BASIC_USER")
                        .requestMatchers(HttpMethod.DELETE,
                                "/comments/*",
                                "/comments/action/like/*"
                        ).hasAuthority("BASIC_USER")
                        .requestMatchers(
                                HttpMethod.GET,
                                "/comments/main",
                                "/comments/reply"
                        ).permitAll()
                        .requestMatchers(HttpMethod.GET, "/api*")
                        .permitAll()
                ).addFilterBefore(new ServiceInterceptor(tokenParser), UsernamePasswordAuthenticationFilter.class)
                /* disable default session management */
                .sessionManagement(smc -> smc.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
                /* disable default csrf protection */
                .csrf(AbstractHttpConfigurer::disable);
        return http.build();
    }
}
