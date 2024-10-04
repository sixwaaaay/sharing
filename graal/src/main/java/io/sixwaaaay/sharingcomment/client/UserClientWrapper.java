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

package io.sixwaaaay.sharingcomment.client;

import io.github.resilience4j.circuitbreaker.CircuitBreaker;
import io.github.resilience4j.circuitbreaker.CircuitBreakerConfig;
import io.github.resilience4j.decorators.Decorators;
import io.github.resilience4j.retry.Retry;
import io.github.resilience4j.retry.RetryConfig;
import io.sixwaaaay.sharingcomment.domain.User;
import org.springframework.web.client.HttpClientErrorException;

import java.time.Duration;
import java.util.Collection;
import java.util.List;

public class UserClientWrapper implements UserClient {
    private final UserClient userClient;

    private final CircuitBreaker circuitBreaker;

    private final Retry retry;

    public UserClientWrapper(UserClient userClient) {

        this.userClient = userClient;
        var name = "userClient";
        var retryConfig = RetryConfig
                .custom()
                .maxAttempts(3)
                .waitDuration(Duration.ofMillis(1000))
                .ignoreExceptions(
                        HttpClientErrorException.NotFound.class,
                        HttpClientErrorException.Forbidden.class,
                        HttpClientErrorException.MethodNotAllowed.class,
                        HttpClientErrorException.Unauthorized.class
                ).build();
        retry = Retry.of(name, retryConfig);
        var circuitBreakerConfig = CircuitBreakerConfig
                .custom()
                .failureRateThreshold(50)
                .ignoreExceptions(
                        HttpClientErrorException.NotFound.class,
                        HttpClientErrorException.Forbidden.class,
                        HttpClientErrorException.MethodNotAllowed.class,
                        HttpClientErrorException.Unauthorized.class
                )
                .waitDurationInOpenState(Duration.ofMillis(1000))
                .permittedNumberOfCallsInHalfOpenState(2)
                .slidingWindowSize(2)
                .build();
        circuitBreaker = CircuitBreaker.of(name, circuitBreakerConfig);
    }

    @Override
    public User getUser(long id, String token) {
        var getUserReplySupplier = Decorators
                .ofSupplier(() -> userClient.getUser(id, token))
                .withCircuitBreaker(circuitBreaker)
                .withRetry(retry)
                .decorate();
        return getUserReplySupplier.get();
    }


    @Override
    public List<User> getManyUser(Collection<Long> ids, String token) {
        var getManyUserReplySupplier = Decorators
                .ofSupplier(() -> userClient.getManyUser(ids, token))
                .withCircuitBreaker(circuitBreaker)
                .withRetry(retry)
                .decorate();
        return getManyUserReplySupplier.get();
    }
}
