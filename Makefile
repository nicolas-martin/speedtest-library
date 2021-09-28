BINARY:=speedtest-test

BUILDOPTS:=-v -ldflags='-s'
BUILDOPTS_STATIC:=-v -ldflags='-s -w -extldflags=-static $(BUILDINJECT)'

.PHONY: all
all: dev

.PHONY: dev
dev:
	GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build $(BUILDOPTS) -o bin/$(BINARY)-$@.mipsle
