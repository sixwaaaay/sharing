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
import io.sixwaaaay.sharingcomment.transmission.*;

import java.time.Duration;
import java.util.Set;

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

    @Override
    public VoteReply itemAdd(VoteReq req) {
        return voteClient.itemAdd(req);
    }

    @Override
    public VoteReply itemDelete(VoteReq req) {
        return voteClient.itemDelete(req);
    }

    /**
     * wrap the exists method with resilience4j circuit breaker and retry
     * fallback with an empty set if the circuit breaker is open
     */
    @Override
    public VoteExistsReply exists(VoteExistsReq req) {
        var existsReplySupplier = Decorators.ofSupplier(() -> voteClient.exists(req))
                .withCircuitBreaker(circuitBreaker)
                .withRetry(retry)
                .withFallback(this::fallback)
                .decorate();
        return existsReplySupplier.get();
    }

    @Override
    public ScanVotedReply scan(ScanVotedReq req) {
        return voteClient.scan(req);
    }

    private VoteExistsReply fallback(Throwable e) {
        var reply = new VoteExistsReply();
        reply.setExists(Set.of());
        return reply;
    }
}
