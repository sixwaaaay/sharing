# Users

[![CI](https://github.com/sixwaaaay/sharing/actions/workflows/users.yaml/badge.svg)](https://github.com/sixwaaaay/sharing/actions/workflows/users.yaml)

[![Container Image Size](https://img.shields.io/docker/image-size/sixwaaaay/shauser/latest)](https://hub.docker.com/r/sixwaaaay/shauser)

[![Docker Pulls](https://img.shields.io/docker/pulls/sixwaaaay/shauser)](https://hub.docker.com/r/sixwaaaay/shauser)

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/sixwaaaay/shauser?sort=semver)](https://hub.docker.com/r/sixwaaaay/shauser)

[![Codcov](https://codecov.io/gh/sixwaaaay/sharing/branch/main/graph/badge.svg)](https://codecov.io/gh/sixwaaaay/sharing)

the users module backend.

## Technologies

| Name        | Usage                |
| ----------- | -------------------- |
| gRPC | gRPC endpoints |
| echo | http endpoints |
| GORM & MySQL | Database |
| GitHub OAuth2 | OAuth2 login |
| testify & go cover | Testing & Coverage |
| Docker | Containerization |
| Compose | Automation |
| GitHub Actions | CI/CD |
| OpenTelemetry-Go | Tracing & Metrics |

## Details

The users module backend provides **both gRPC and http APIs** for users.
gRPC serializes the data with protobuf, and http serializes the data with json.

**GORM** is used to access the MySQL database and change the data source.

**GitHub OAuth** is used to provide the OAuth login feature, which is a third-party login feature so that users can log in with their GitHub accounts without registering.

**Testify** is used to simplify the testing, and the built-in **go cover** is used to collect the coverage.

**Docker** is used to containerize the application, and **Compose** is used to automate the development environment.

**GitHub Actions** is used to automate the CI/CD process.

**OpenTelemetry-Go** is used to collect and export the tracing and metrics data.

## User Features

- Sign Up and Sign In
- OAuth Login with GitHub
- Get User Info
- Update User Info
- Follow and Unfollow Users
- Get Followers and Followings List by Page

## Service Features

- gRPC and http APIs
- Read and write separation
- Database connection pool
- Cached Precompiled Statements
- OAuth login
- Tracing and metrics by OpenTelemetry protocol
