GO_PKG_DIRS  := $(subst $(shell go list -e -m),.,$(shell go list ./... | grep -v /vendor | grep -v /policy-server ))

all:
	$(info GOPI1 $(shell go list ./... | grep -v /vendor | grep -v /policy-server ))
	$(info GOPI2 $(GO_PKG_DIRS))
	go build -ldflags="-s -w" -o server $(GO_PKG_DIRS)

fmt:
	gofmt -s -w $(GO_PKG_DIRS)

lint:
	golangci-lint run -v $(GO_PKG_DIRS)

clean:
	rm -f server
