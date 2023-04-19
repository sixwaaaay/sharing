FROM golang:1.20-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED 0
#ENV GOPROXY https://goproxy.cn,direct
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/account

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
WORKDIR /app
COPY --from=builder /app/account /app/account
# configs
COPY configs /app/configs
EXPOSE 8080
CMD ["./account"]
