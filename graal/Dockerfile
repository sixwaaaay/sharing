FROM debian:bookworm-slim
ARG APP_NAME=sharing-comment
WORKDIR /app
COPY ["./target/${APP_NAME}","/app/app"]
ENTRYPOINT ["/app/app"]