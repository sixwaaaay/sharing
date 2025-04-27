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

import com.zaxxer.hikari.HikariDataSource;
import io.sixwaaaay.sharingcomment.util.DbContextEnum;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.autoconfigure.jdbc.DataSourceProperties;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Lazy;
import org.springframework.context.annotation.Primary;
import org.springframework.data.jdbc.core.convert.*;
import org.springframework.data.jdbc.core.mapping.JdbcMappingContext;
import org.springframework.data.relational.core.dialect.Dialect;
import org.springframework.data.relational.core.dialect.PostgresDialect;
import org.springframework.data.relational.core.mapping.DefaultNamingStrategy;
import org.springframework.data.relational.core.mapping.NamingStrategy;
import org.springframework.jdbc.core.namedparam.NamedParameterJdbcOperations;
import org.springframework.jdbc.core.namedparam.NamedParameterJdbcTemplate;
import org.springframework.jdbc.datasource.DataSourceTransactionManager;
import org.springframework.transaction.TransactionManager;

import javax.sql.DataSource;
import java.util.HashMap;
import java.util.Optional;

/**
 * This class is responsible for configuring the data sources for the application.
 * It sets up both read and write data sources, and configures the routing between them.
 * It also sets up the JDBC operations, transaction manager, dialect, custom conversions, mapping context, and converter.
 * The data sources are configured using the HikariCP library, which is a high-performance JDBC connection pool.
 * The routing between the data sources is done using a custom RoutingDataSource.
 * The JDBC operations are set up using the NamedParameterJdbcTemplate class from Spring JDBC.
 * The transaction manager is set up using the DataSourceTransactionManager class from Spring JDBC.
 */
@Configuration
public class DataSourceConfig {

    /**
     * RoutingDataSource is a custom DataSource that can switch between different DataSources
     *
     * @param readDataSource  readDataSource
     * @param writeDataSource writeDataSource
     * @return routingDataSource
     */
    @Bean
    @Primary // make this the primary DataSource
    public DataSource routingDataSource(@Qualifier("readDataSource") DataSource readDataSource, @Qualifier("writeDataSource") DataSource writeDataSource) {
        var targetDataSources = new HashMap<>();
        targetDataSources.put(DbContextEnum.READ, readDataSource);
        targetDataSources.put(DbContextEnum.WRITE, writeDataSource);

        var routingDataSource = new RoutingDataSource();
        routingDataSource.setTargetDataSources(targetDataSources);
        routingDataSource.setDefaultTargetDataSource(writeDataSource);

        return routingDataSource;
    }

    // the write data source properties
    @Bean
    @Primary
    @ConfigurationProperties("spring.datasource")
    public DataSourceProperties writeDataSourceProperties() {
        return new DataSourceProperties();
    }

    // the write data source
    @Bean
    public HikariDataSource writeDataSource(@Qualifier("writeDataSourceProperties") DataSourceProperties properties) {
        return properties.initializeDataSourceBuilder().type(HikariDataSource.class).build();
    }


    // the read data source properties
    @Bean
    @ConfigurationProperties("spring.replica-datasource")
    public DataSourceProperties readDataSourceProperties() {
        return new DataSourceProperties();
    }

    // the read data source
    @Bean
    public HikariDataSource readDataSource(@Qualifier("readDataSourceProperties") DataSourceProperties properties) {
        return properties.initializeDataSourceBuilder().type(HikariDataSource.class).build();
    }

    // the JDBC operations
    //
    @Bean
    @Primary
    NamedParameterJdbcOperations namedParameterJdbcOperations(DataSource dataSource) {
        return new NamedParameterJdbcTemplate(dataSource);
    }

    // the transaction manager
    @Bean
    TransactionManager transactionManager(DataSource dataSource) {
        return new DataSourceTransactionManager(dataSource);
    }

    // sql dialect
    @Bean
    Dialect jdbcDialect() {
        return PostgresDialect.INSTANCE;
    }

    // custom conversions
    @Bean
    JdbcCustomConversions customConversions() {
        return new JdbcCustomConversions();
    }

    // mapping context
    @Bean
    JdbcMappingContext jdbcMappingContext(Optional<NamingStrategy> namingStrategy, JdbcCustomConversions customConversions) {
        var mappingContext = new JdbcMappingContext(namingStrategy.orElse(DefaultNamingStrategy.INSTANCE));
        mappingContext.setSimpleTypeHolder(customConversions.getSimpleTypeHolder());
        return mappingContext;
    }

    // converter
    @Bean
    JdbcConverter jdbcConverter(JdbcMappingContext context, NamedParameterJdbcOperations operations, @Lazy RelationResolver relationResolver, JdbcCustomConversions conversions, Dialect dialect) {
        // return new BasicJdbcConverter(context, relationResolver, conversions, jdbcTypeFactory, dialect.getIdentifierProcessing());
        return new MappingJdbcConverter(context, relationResolver);
    }
}