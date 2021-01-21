NAME := crypto
DOCKER_REPOSITORY := json0
DOCKER_IMAGE_NAME := $(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT := $(shell git describe --dirty --always)
SOURCE_FILES := $(shell find * -name '*.go')
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')
EXTRA_RUN_ARGS?=

crypto: $(SOURCE_FILES)
	export GO111MODULE=on
	go mod download
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/jasonbirchall/crypto/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./crypto ./main.go

test:
	go test -v ./...

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

release:
	git tag $(VERSION)
	git push origin $(VERSION)

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
  go get github.com/jasonbirchall/crypto

.PHONY: test
