VERSION 0.8

FROM golang:1.23.2-alpine

# install gcc dependencies into alpine for CGO
RUN apk --no-cache add git ca-certificates gcc musl-dev libc-dev binutils-gold curl openssh

# install docker tools
# https://docs.docker.com/engine/install/debian/
RUN apk add --update --no-cache docker

# install linter
# binary will be $(go env GOPATH)/bin/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1
RUN ls -la $(which golangci-lint)

RUN go install github.com/ory/go-acc@latest

# install vektra/mockery
RUN go install github.com/vektra/mockery/v2@v2.43.2


all:
    BUILD +lint
    BUILD +test

code:
    # download deps
    COPY go.mod go.sum .
    RUN go mod download -x
    # copy code
    COPY --dir . ./

vendor:
    FROM +code
    RUN go mod vendor
    SAVE ARTIFACT /app /files

lint:
    FROM +vendor

    COPY .golangci.yml ./

    # Runs golangci-lint with settings:
    RUN golangci-lint run --timeout 10m --skip-dirs-use-default

test:
    FROM +vendor

    RUN go test -mod=vendor ./... -race -coverprofile=coverage.out -covermode=atomic -coverpkg=./...

    SAVE ARTIFACT coverage.out AS LOCAL coverage.out

