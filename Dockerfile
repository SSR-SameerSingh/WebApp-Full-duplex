FROM golang:alpine AS build-env

WORKDIR /go-work/src
ADD . /go-work/src/WebSockets/WebApp-Full-duplex
RUN cd /go-work/src/WebSockets/WebApp-Full-duplex && go build -o main

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*

WORKDIR /app

COPY --from=build-env /go-work/src/WebSockets/WebApp-Full-duplex /app/static

EXPOSE 8080
ENTRYPOINT [ “./main” ]