package io.sixwaaaay.sharingcomment.config;


import org.springframework.boot.context.properties.ConfigurationProperties;


@ConfigurationProperties(prefix = "service")
public record ServiceConfig(Item vote, Item user) {
    public record Item(boolean enabled, String baseUrl) {
    }
}
