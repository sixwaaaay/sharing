# Content

[![CI](https://github.com/sixwaaaay/sharing/actions/workflows/content.yaml/badge.svg)](https://github.com/sixwaaaay/sharing/actions/workflows/content.yaml)

[![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/content/latest)](https://hub.docker.com/r/sixwaaaay/content)

[![Docker Pulls](https://img.shields.io/docker/pulls/sixwaaaay/content)](https://hub.docker.com/r/sixwaaaay/content)

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/content?sort=semver)](https://hub.docker.com/r/sixwaaaay/content)

[![Codcov](https://codecov.io/gh/sixwaaaay/sharing/branch/main/graph/badge.svg)](https://codecov.io/gh/sixwaaaay/sharing)

The content module backend.

## Technologies

| Name        | Usage                |
| ----------- | -------------------- |
| Minimal API | Application Backbone  Framework   |
| Dapper.AOT | Database Access |
| MySQL | Database |
| NativeAOT | AOT Compilation |
| Grpc Client | Grpc Call |
| Bearer Token | Security |
| Prometheus | Metrics |
| OpenTelemetry | Tracing |
| XUnit | Testing |
| Coverlet | Coverage |
| Docker | Containerization |
| Compose | Automation |
| GitHub Actions | CI/CD |

## Details

The content module backend is a Minimal API application that provides HTTP APIs for content.
With the help of **Dapper.AOT**, the application can access the MySQL database with simplified object mapping without reflection.
The application is compiled to a native image by **NativeAOT**, so that the application can start faster, consume less memory and have a smaller disk footprint.
**Grpc Client** is used to call the user module backend to get the user information.
**Prometheus** is used to collect the metrics, and **OpenTelemetry** is used to collect the traces.
**Docker** is used to containerize the application, and **Compose** is used to automate the development environment.
**GitHub Actions** is used to automate the CI/CD process, and **XUnit** is used to test the application with **Coverlet** for coverage.

## User Features

- Submit a video contents
- Read video contents by page
- Read video contents of a user by page
- Like a video content
- Dislike a video content
- Chat with a user
- Query Chat history with a user
- Video Information probing

## Service Features

- MySQL Index Optimization for video contents and chat history query performance
- Keyset Pagination for video contents and chat history query performance
- AOT Compilation for less memory and faster startup
- Docker Containerization for packaging
- Compose Automation for development environment
- GitHub Actions for CI/CD
- XUnit for testing
- Coverlet for coverage
