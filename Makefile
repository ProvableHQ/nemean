.PHONY: all test clean

GO_PKG_LIST := $(shell go list ./...)

.PHONY: vet
vet:
	go vet ${GO_PKG_LIST}

.PHONY: gofmt
gofmt:
	test -z "$(shell gofmt -s -l $(shell find ./* -name '*.go' ! -name '*_test.go' -type f -print) | tee /dev/stderr)"

.PHONY: test
test:
	GOFLAGS=-mod=vendor go test -short ${GO_PKG_LIST}

.PHONY: staticcheck
staticcheck:
	GOFLAGS=-mod=vendor staticcheck ${GO_PKG_LIST}

.PHONY: lint
lint:
	golint -set_exit_status ${GO_PKG_LIST}

.PHONY: build
build:
	cd aleo && cargo build --release
	cp aleo/target/release/libaleo.so /usr/lib/libaleo.so
	cp aleo/aleo.h /usr/include/aleo.h
	go build -ldflags="-r aleo" -o nemean cmd/nemean/*.go