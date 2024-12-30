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

import java.time.Duration;
import java.util.List;

public class VoteClientWrapper implements VoteClient {


    private final VoteClient voteClient;

    public VoteClientWrapper(VoteClient voteClient) {
        this.voteClient = voteClient;
    }

    private final String name = "voteClient";

    private final CircuitBreaker circuitBreaker = CircuitBreaker.of(name, CircuitBreakerConfig.custom()
            .failureRateThreshold(50)
            .waitDurationInOpenState(Duration.ofMillis(2000))
            .permittedNumberOfCallsInHalfOpenState(2)
            .slidingWindowSize(2)
            .build()
    );

    private final Retry retry = Retry.of(name, RetryConfig.custom()
            .maxAttempts(3)
            .waitDuration(Duration.ofMillis(1000))
            .build()
    );
    /**
     * wrap the queryInLikes method with resilience4j circuit breaker and retry
     * @param commentIds list of object ids
     * @param token user token
     * @return list of object ids that user liked
     */
    @Override
    public List<Long> queryInLikes(List<Long> commentIds, String token) {
        var queryInLikesSupplier = Decorators.ofSupplier(() -> voteClient.queryInLikes(commentIds, token))
                .withCircuitBreaker(circuitBreaker)
                .withRetry(retry)
                .withFallback((e) -> List.of())
                .decorate();
        return queryInLikesSupplier.get();
    }
}
