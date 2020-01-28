# --------------------------------------
# Build Stage
# --------------------------------------
FROM golang:latest
RUN mkdir /app 
ENV GOPATH /go
ADD . /go/src/github.com/SSR-SameerSingh/WebApp-Full-duplex
WORKDIR /go/src/github.com/SSR-SameerSingh/WebApp-Full-duplex
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

# --------------------------------------
# Production Container
# --------------------------------------
FROM scratch
COPY --from=build_stage /go/src/github.com/SSR-SameerSingh/WebApp-Full-duplex/WebApp-Full-duplex
CMD [/WebApp-Full-duplex]