# Comments

[![CI](https://github.com/sixwaaaay/sharing/actions/workflows/comment.yaml/badge.svg)](https://github.com/sixwaaaay/sharing/actions/workflows/comment.yaml)

[![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/sharing-comment/latest)](https://hub.docker.com/r/sixwaaaay/sharing-comment)

[![Docker Pulls](https://img.shields.io/docker/pulls/sixwaaaay/sharing-comment)](https://hub.docker.com/r/sixwaaaay/sharing-comment)

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/sharing-comment?sort=semver)](https://hub.docker.com/r/sixwaaaay/sharing-comment)

[![Codcov](https://codecov.io/gh/sixwaaaay/sharing/branch/main/graph/badge.svg)](https://codecov.io/gh/sixwaaaay/sharing)

the comments module backend.

## Technologies

| Name        | Usage                |
| ----------- | -------------------- |
| Spring Boot | Application Backbone  Framework   |
| Spring Cache & Redis | Caching |
| Spring Data JDBC & MySQL | Database |
| Resilience4j | Circuit Breaker & Retry |
| Micrometer | Observability |
| Prometheus | Metrics |
| OpenTelemetry | Tracing |
| Bearer Token | Security |
| GraalVM JDK 21 | AOT Compilation |
| Docker | Containerization |
| Compose | Automation |
| GitHub Actions | CI/CD |
| JUnit 5 | Testing |
| Jacoco | Coverage |

## Details

The comments module backend is a Spring Boot application that provides HTTP APIs for comments.

Thanks to the Spring Boot framework, the application is easy to develop.
**The minimal JDK version is 21**, so that **virtual threads** can be used to improve the performance.

**Spring Data JDBC** and **Spring Data Relational** are used to access the MySQL database and generate the SQL queries automatically.

**Spring Cache** is used to cache the comments data, and **Redis** is used as the cache store. The cache serialization is done by **Jackson**.

**Resilience4j** is used to provide the circuit breaker and retry features, so that the application can be more resilient and degrade when necessary.

**Micrometer** is used to collect the metrics, and **Prometheus** is used to store the metrics. **OpenTelemetry** is used to collect the traces.

**GraalVM JDK 21** is used to compile the application to a native image, so that the application can start faster, consume less memory and have a smaller disk footprint.

**Docker** is used to containerize the application, and **Compose** is used to automate the development environment.

**GitHub Actions** is used to automate the CI/CD process, and **JUnit 5** is used to test the application with **Jacoco** for coverage.

## User Features

- Submit a comment
- Read comments by page
- Read sub-comments by page
- Reply to a comment or sub-comment
- Like a comment or sub-comment
- Dislike a comment or sub-comment

## Service Features

- Dynamic DataSource Routing for Read/Write Separation
- MySQL Database Index for Comment Query
- Redis Cache for Comment Data
- keyset Pagination for Comment Query
- AOT Compilation for GraalVM Native Image
- Docker Containerization for packaging
- Compose Automation for development environment
- GitHub Actions for CI/CD
- JUnit 5 & Jacoco for testing and coverage

## Service Dependencies

The comments module backend depends on the following services:

- MySQL Database for storing the comments data
- Redis for caching, this is optional,which can be disabled by configuration
- OpenTelemetry Collector for collecting the traces, this is optional, which can be disabled by configuration
- Users Service with http API for user information, this is optional, which can be disabled by configuration
- Vote Service with http API for voting, this is optional, which can be disabled by configuration
