/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.config;

import io.sixwaaaay.sharingcomment.client.UserClient;
import io.sixwaaaay.sharingcomment.client.VoteClient;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestClient;
import org.springframework.web.client.support.RestClientAdapter;
import org.springframework.web.service.invoker.HttpServiceProxyFactory;

@Configuration
public class Config {
    @Bean
    VoteClient repositoryService(
            @Value("${service.vote.base-url}") String baseUrl,
            RestClient.Builder restClientBuilder
    ) {
        return HttpServiceProxyFactory.builderFor(
                RestClientAdapter.create(restClientBuilder.baseUrl(baseUrl).build())
        ).build().createClient(VoteClient.class);
    }

    @Bean
    UserClient userService(
            @Value("${service.user.base-url}") String baseUrl,
            RestClient.Builder restClientBuilder
    ) {
        return HttpServiceProxyFactory.builderFor(
                RestClientAdapter.create(restClientBuilder.baseUrl(baseUrl).build())
        ).build().createClient(UserClient.class);
    }

}
