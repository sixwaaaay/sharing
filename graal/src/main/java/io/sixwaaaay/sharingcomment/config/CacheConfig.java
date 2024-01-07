/*
 * Copyright (c)  sixwaaaay
 * All rights reserved.
 */

package io.sixwaaaay.sharingcomment.config;


import com.fasterxml.jackson.databind.ObjectMapper;
import io.sixwaaaay.sharingcomment.domain.Comment;
import org.springframework.boot.autoconfigure.cache.RedisCacheManagerBuilderCustomizer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.cache.RedisCacheConfiguration;
import org.springframework.data.redis.serializer.Jackson2JsonRedisSerializer;
import org.springframework.data.redis.serializer.RedisSerializationContext;

import java.time.Duration;
import java.util.List;

@Configuration(proxyBeanMethods = false)
public class CacheConfig {
    @Bean
    public RedisCacheManagerBuilderCustomizer redisCacheManagerBuilderCustomizer(ObjectMapper objectMapper) {
        var type = objectMapper.getTypeFactory().constructCollectionType(List.class, Comment.class);

        var redisCacheConfiguration = RedisCacheConfiguration.defaultCacheConfig()
                .entryTtl(Duration.ofMinutes(6))
                .disableCachingNullValues()
                .serializeValuesWith(
                        RedisSerializationContext.SerializationPair.fromSerializer(
                                new Jackson2JsonRedisSerializer<List<Comment>>(objectMapper, type)
                        )
                );
        return (builder) -> builder
                .withCacheConfiguration("comments-main", redisCacheConfiguration)
                .withCacheConfiguration("comments-reply", redisCacheConfiguration);
    }
}
