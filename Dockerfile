FROM golang:1.3

ADD dockerhub-webhook-listener /go/src/github.com/cpuguy83/dockerhub-webhook-listener
WORKDIR /go/src/github.com/cpuguy83/dockerhub-webhook-listener/hub-listener
RUN go get && go build
ADD entrypoint.sh /entrypoint.sh
ADD config.ini /config.ini
EXPOSE 80
CMD exec /entrypoint.sh
