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

import io.sixwaaaay.sharingcomment.client.UserClient;
import io.sixwaaaay.sharingcomment.client.UserClientWrapper;
import io.sixwaaaay.sharingcomment.client.VoteClient;
import io.sixwaaaay.sharingcomment.client.VoteClientWrapper;
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
        var client = createService(VoteClient.class, baseUrl, restClientBuilder);
        return new VoteClientWrapper(client);
    }

    @Bean
    UserClient userService(
            @Value("${service.user.base-url}") String baseUrl,
            RestClient.Builder restClientBuilder
    ) {
        var client = createService(UserClient.class, baseUrl, restClientBuilder);
        return new UserClientWrapper(client);
    }

    private <T> T createService(Class<T> clazz, String baseUrl, RestClient.Builder restClientBuilder) {
        var restClient = restClientBuilder.baseUrl(baseUrl).build();
        var adapter = RestClientAdapter.create(restClient);
        var proxyFactory = HttpServiceProxyFactory.builderFor(adapter);
        return proxyFactory.build().createClient(clazz);
    }
}
