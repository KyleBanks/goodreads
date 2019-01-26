test:
	go test ./... -vet all
.PHONY: test

deps:
	GO111MODULE=on GOFLAGS=-mod=vendor go mod tidy
	GO111MODULE=on GOFLAGS=-mod=vendor go mod vendor
.PHONY: deps
