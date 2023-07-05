FROM golang:1.18-alpine

WORKDIR /app
COPY . .
RUN cp config.docker.yaml config.yaml
RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go mod download && go build -o main && go build -o fake ./debugs
# 先执行fake，再执行main
CMD ["./docker-entrypoint.sh"]

#FROM alpine:latest
#WORKDIR /root
#COPY --from=0 /app/main ./main
#COPY --from=0 /app/config.docker.yaml ./config.yaml
#CMD ["./main"]