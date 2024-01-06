/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.config;

import io.sixwaaaay.sharingcomment.client.UserClient;
import io.sixwaaaay.sharingcomment.client.VoteClient;
import io.sixwaaaay.sharingcomment.request.Principal;
import org.springframework.aot.hint.annotation.RegisterReflectionForBinding;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestClient;
import org.springframework.web.client.support.RestClientAdapter;
import org.springframework.web.service.invoker.HttpServiceProxyFactory;

import java.util.ArrayList;

@Configuration
@RegisterReflectionForBinding({Principal.class, ArrayList.class})
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
