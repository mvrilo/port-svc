.PHONY: all test build image run-image

all: test build

test:
	go test -race ./...

build:
	CGO_ENABLED=0 go build -o port-svc ./cmd/port-svc/*.go

image:
	docker build . -t mvrilo/port-svc

run-image:
	docker run --rm -it mvrilo/port-svc
