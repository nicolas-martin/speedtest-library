BINARY:=speedtest-test

BUILDOPTS:=-v -ldflags='-s'
BUILDOPTS_STATIC:=-v -ldflags='-s -w -extldflags=-static'

.PHONY: all
all: dev

.PHONY: dev
dev:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build $(BUILDOPTS) -o bin/$(BINARY)-$@.mipsle

.PHONY: linux
linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(BUILDOPTS_STATIC) -o bin/$(BINARY)-router.linux.amd64
