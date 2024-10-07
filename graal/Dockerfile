FROM debian:bookworm-slim

WORKDIR /app
COPY ["./target/sharing-comment", "/app/sharing-comment"]

ENTRYPOINT ["/app/sharing-comment"]