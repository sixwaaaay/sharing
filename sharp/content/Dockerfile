#
# Copyright (c) 2023-2024 sixwaaaay.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# 
#
FROM alpine:3.20 AS base
WORKDIR /app
RUN adduser -u 1000 --gecos "" --disabled-password appuser && chown -R appuser /app
USER appuser
EXPOSE 5000

FROM mcr.microsoft.com/dotnet/sdk:9.0.100-alpine3.20 AS build
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
