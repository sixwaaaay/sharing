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

package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Config struct {
	ListenOn     string // ListenOn is the address gRPC server listening on
	HTTP         string // HTTP is the address of the HTTP server
	DefaultLimit int32  // DefaultLimit is the default limit of the query
	MaxLimit     int32  // MaxLimit is the max limit of the query
	Secret       string // Secret is the secret used to sign the JWT
	TTL          string // TTL is the expiration duration of JWT
	MySQL        struct {
		DSN      string   // MySQLDSN is the DSN of the MySQL database
		Replicas []string // MySQLDSN is the replicas database dsn of the MySQL database
	} // MySQL is the configuration for the MySQL database
	Otel struct {
		Enabled     bool   `yaml:"enabled"`     // OtelEnabled is the flag to enable OpenTelemetry
		Endpoint    string `yaml:"endpoint"`    // OtelEndpoint is the endpoint of the OpenTelemetry collector
		Service     string `yaml:"service"`     // OtelService is the service name
		Version     string `yaml:"version"`     // OtelVersion is the version of service
		Environment string `yaml:"environment"` // OtelEnvironment is the environment of service
	} `yaml:"otel"` // Otel is the configuration for the OpenTelemetry collector
	Cache struct {
		Enabled bool   `yaml:"enabled"` // CacheEnabled is the flag to enable cache
		TTL     string `yaml:"ttl"`     // CacheTTL is the expiration duration of cache
	}
	Redis redis.UniversalOptions
	Oauth oauth2.Config
}

// NewConfig parses the config file and returns a Config struct
func NewConfig(p string) (Config, error) {
	viper.SetConfigFile(p)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
