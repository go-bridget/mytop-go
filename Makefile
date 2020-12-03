.PHONY: all drone build test

all: drone

drone:
	@drone exec --exclude notify --registry=http://$(shell jq -r '.auths["registry.dev.rtvslo.si"].auth' /root/.docker/config.json | base64 --decode)@registry.dev.rtvslo.si

build:
	mkdir -p build
	go build -o build/ ./cmd/...

test:
	go mod tidy
	go fmt ./...
	go test ./... -v --tags="integration debug"
