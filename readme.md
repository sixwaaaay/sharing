# sharing

Sharing Has An Relly Interesting Name

## Description

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsixwaaaay%2Fsharing.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsixwaaaay%2Fsharing?ref=badge_shield)
[![Apache 2.0](https://img.shields.io/github/license/sixwaaaay/sharing)](https://github.com/sixwaaaay/sharing/blob/main/LICENSE)

| name    | CI                                                                                                                                                           | coverage                                                                                                                                 | size                                                                                                                                                          | version                                                                                                                                                              |
|---------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| users   | [![CI](https://github.com/sixwaaaay/sharing/actions/workflows/users.yaml/badge.svg)](https://github.com/sixwaaaay/sharing/actions/workflows/users.yaml)      | [![codecov](https://codecov.io/gh/sixwaaaay/sharing/branch/main/graph/badge.svg?flag=users)](https://codecov.io/gh/sixwaaaay/sharing)    | [![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/shauser?sort=semver)](https://hub.docker.com/r/sixwaaaay/shauser)                 | [![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/shauser?sort=semver)](https://hub.docker.com/r/sixwaaaay/shauser)                 |
| content | [![CI](https://github.com/sixwaaaay/sharing/actions/workflows/content.yaml/badge.svg)](https://github.com/sixwaaaaay/sharing/actions/workflows/content.yaml) | [![codecov](https://codecov.io/gh/sixwaaaay/sharing/branch/main/graph/badge.svg?flag=content)](https://codecov.io/gh/sixwaaaay/sharing)  | [![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/content?sort=semver)](https://hub.docker.com/r/sixwaaaay/content)                 | [![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/content?sort=semver)](https://hub.docker.com/r/sixwaaaay/content)                 |
| comment | [![CI](https://github.com/sixwaaaay/sharing/actions/workflows/comment.yaml/badge.svg)](https://github.com/sixwaaaaay/sharing/actions/workflows/comment.yaml) | [![codecov](https://codecov.io/gh/sixwaaaay/sharing/branch/main/graph/badge.svg?flag=comments)](https://codecov.io/gh/sixwaaaay/sharing) | [![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/sharing-comment?sort=semver)](https://hub.docker.com/r/sixwaaaay/sharing-comment) | [![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/sharing-comment?sort=semver)](https://hub.docker.com/r/sixwaaaay/sharing-comment) |

sharing is a content social platform backend. OpenAPI spec can be previewed
at [users](https://cdn.quzhao.me/apidoc.html?url=/end.json), [content](https://cdn.quzhao.me/apidoc.html?url=/content.json), [comment](https://cdn.quzhao.me/apidoc.html?url=/comments.yaml).
`openapi-typescript-codegen` is recommended for typescript client code generation.

## Requirements

A Github OAuth 2.0 Key is required. You can create a new OAuth App in your Github account.

A Mailgun API Key is required. You can create a new API Key in your Mailgun account.

A MySQL(8.0+) Cluster. MySQL Group Replication Cluster deployed by MySQL Operator is recommended. GTID mode is required.

A PostgreSQL(16.0+) Cluster. Patroni high availability (HA) PostgreSQL is recommended.

A S3 Compatible Object Storage Server. AWS S3 is recommended.

A Redis(5.0+) Cluster.

A RisingWave(1.10+) Cluster. deployed by RisingWave Operator is recommended.

A OpenTelemetry Collector(0.33+) Cluster. deployed by OpenTelemetry Operator is recommended.

A OpenTelemetry Data Storage. Grafana Cloud is recommended.

A ElasticSearch(7.0+) Cluster.

## Installation

the application container images are build by GitHub Actions, you can pull the images
from [Docker Hub](https://hub.docker.com/u/sixwaaaay).

the sequence of deployment
is [users](https://hub.docker.com/r/sixwaaaay/shauser), [content](https://hub.docker.com/r/sixwaaaay/content), [comment](https://hub.docker.com/r/sixwaaaay/sharing-comment).
note that the `users` service is required by `content` and `comment` service.

for `content` and `comment` service, configuration can be override by environment variables.
`users` service need a configuration file `config.yaml` in configs directory of the container.

just deploy them as a Kubernetes Deployment. Horizontal Pod Autoscaler(HPA) is recommended.

## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsixwaaaay%2Fsharing.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsixwaaaay%2Fsharing?ref=badge_large)
