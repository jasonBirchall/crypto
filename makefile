SOURCE_FILES := $(shell find * -name '*.go')

crypto-tracker: $(SOURCE_FILES)
	export GO111MODULE=on
	go mod download
	go build -o crypto ./main.go

test:
	go test ./...

fmt:
	go fmt ./...

.PHONY: test
