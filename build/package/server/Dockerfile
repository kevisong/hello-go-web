FROM alpine

RUN mkdir -p /data/hello-go-web
WORKDIR /data/hello-go-web
ADD ./cmd/server/server .

EXPOSE 80
EXPOSE 8080

CMD ['/data/hello-go-web/server', "--config=/data/hello-go-web/config/config.yml"]
