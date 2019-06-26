FROM golang:1.12.6-alpine3.10 AS build-env

#We don't want to run the cgo resolver
ENV CGO_ENABLED 0

# Allow Go to retrive the dependencies for the build step
RUN apk add --no-cache git

WORKDIR /myproject/
ADD . /myproject/

# Compile the binary
RUN go build -o /myproject/demo .

WORKDIR /go/src/
RUN go get github.com/go-delve/delve/cmd/dlv

# final stage
FROM alpine:3.10

WORKDIR /
COPY --from=build-env /myproject/demo /
COPY --from=build-env /go/bin/dlv /

EXPOSE 8080 40000

CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "--log", "--log-output=debugger,rpc", "exec", "/demo"]
