FROM alpine:3.19 AS base
WORKDIR /app
RUN adduser -u 1000 --gecos "" --disabled-password appuser && chown -R appuser /app
USER appuser
EXPOSE 5000

FROM mcr.microsoft.com/dotnet/sdk:8.0-alpine3.18 AS build
RUN apk update && apk upgrade && apk add --no-cache clang build-base zlib-dev grpc-plugins
WORKDIR /src
ARG RUNTIME_ID=linux-musl-x64
COPY ["./content.csproj", "./"]
RUN dotnet restore content.csproj -r $RUNTIME_ID
COPY . .

RUN PROTOBUF_PROTOC=/usr/bin/protoc  \
    GRPC_PROTOC_PLUGIN=/usr/bin/grpc_csharp_plugin  \
    dotnet publish content.csproj  -c Release --no-restore -r $RUNTIME_ID -o /app/publish /p:DebugType=None /p:DebugSymbols=false

FROM base AS final
WORKDIR /app
ENV ASPNETCORE_URLS=http://0.0.0.0:5000
COPY --chown=appuser --from=build /app/publish .
ENTRYPOINT ["./content"]
