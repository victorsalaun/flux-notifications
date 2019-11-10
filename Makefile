GOOS ?= $(shell go env GOOS)
GOARCH = amd64

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=1 go build -o flux-notifications .
