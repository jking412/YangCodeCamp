FROM golang:1.18-alpine

WORKDIR /app
COPY . .
RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go mod download && go build -o main

FROM alpine:latest
WORKDIR /root
COPY --from=0 /app/main ./main
COPY --from=0 /app/config.docker.yaml ./config.yaml
CMD ["./main"]