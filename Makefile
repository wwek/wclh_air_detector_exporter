PKG_LIST := $(shell go list ./...)

export PATH := ./bin:$(PATH)
export GO111MODULE := on
export GOPROXY = https://goproxy.cn,direct

.PHONY: setup
setup:
	go mod download
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh

.PHONY: modules
modules:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: build
build:
	go build

.PHONY: buildall
buildall:
	./bin/goreleaser --snapshot --skip-publish --rm-dist

.PHONY: lint
lint:
	./bin/golangci-lint run

.PHONY: test
test: 
	go clean -testcache ${PKG_LIST}
	go test -short --race ${PKG_LIST}

# gofmt and goimports all go files
.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

default: build
