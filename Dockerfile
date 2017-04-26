FROM golang:latest

ADD ./ /go/src/device-listener-go

RUN go get github.com/streadway/amqp
RUN go install device-listener-go

ENTRYPOINT /go/bin/device-listener-go

EXPOSE 7610