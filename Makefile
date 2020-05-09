.PHONY:
	build move-to-bin install test lint dep test-ci latest-release

dep:
	GO111MODULE=on go mod vendor

build:
	@echo "Building the binary..."
	@go build -o devlog .
	@echo "Done."

move-to-bin:
	@echo "Moving binary to /usr/local/bin...requires sudo"
	@sudo cp devlog /usr/local/bin
	@echo "Done."

install:
	make build && make move-to-bin

lint:
	golint pkg/

test:
	go test ./...

test-ci:
	gotestsum --junitfile pkg/results.xml

latest-release:
	curl -sL https://api.github.com/repos/automoto/devlog/releases/latest | jq -r '.assets[].browser_download_url'
