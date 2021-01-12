NAME := crypto-tracker
DOCKER_REPOSITORY := json0
DOCKER_IMAGE_NAME := $(DOCKER_REPOSITORY)/$(NAME)
GIT_COMMIT := $(shell git describe --dirty --always)
SOURCE_FILES := $(shell find * -name '*.go')
VERSION:=$(shell grep 'VERSION' internal/version/version.go | awk '{ print $$4 }' | tr -d '"')
EXTRA_RUN_ARGS?=

crypto-tracker: $(SOURCE_FILES)
	export GO111MODULE=on
	go mod download
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/jasonbirchall/crypto-tracker/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./crypto ./main.go

run-tracker:
	go run main.go track --coin btc,eth

test:
	go test -v ./...

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

release:
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: test
