# sharing-comments

the comments module backend.

[![CI](https://github.com/sixwaaaay/comments/actions/workflows/comment.yaml/badge.svg)](https://github.com/sixwaaaay/comments/actions/workflows/comment.yaml)
[![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/sharing-comment/latest)](https://hub.docker.com/r/sixwaaaay/sharing-comment)
[![Docker Pulls](https://img.shields.io/docker/pulls/sixwaaaay/sharing-comment)](https://hub.docker.com/r/sixwaaaay/sharing-comment)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/sharing-comment?sort=semver)](https://hub.docker.com/r/sixwaaaay/sharing-comment)

the comments module is based on spring boot, spring http interface, spring cache etc.
**virtual thread** and **graalvm native image** are used to improve performance.

## Quick Start

for development start the application with docker-compose:

```shell
docker-compose up
```

## Features

- [x] comment for specified object
- [x] key-based pagination for comments
- [x] comment reply
- [x] comment vote
- [x] opentelemetry tracing
- [x] prometheus metrics