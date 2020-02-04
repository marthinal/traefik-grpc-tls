FROM golang:1.12.5-alpine

ENV CGO_ENABLED 0

WORKDIR /go/src/traefik-grpc-tls/server

ADD ./ /go/src/traefik-grpc-tls

RUN apk add git
RUN go get -v

RUN apk add --no-cache git

RUN go build -o /go/src/traefik-grpc-tls/grpc-server traefik-grpc-tls/server

RUN apk add --no-cache libc6-compat

EXPOSE 5300

CMD ["/go/src/traefik-grpc-tls/grpc-server"]
