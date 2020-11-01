.PHONY:
	build move-to-bin install test lint dep test-ci latest-release

dep:
	GO111MODULE=on go mod vendor

build:
	@echo "Building the binary..."
	@go build -v -ldflags="-X 'main.devlogVersion=dev'" -o devlog .
	@echo "Done."

move-to-bin:
	@echo "Moving binary to /usr/local/bin...requires sudo"
	@sudo cp devlog /usr/local/bin
	@echo "Done."

install:
	make build && make move-to-bin

lint:
	golint -set_exit_status pkg/

test:
	go test ./...

test-ci:
	gotestsum --format testname --junitfile pkg/results.xml -- -coverprofile=cover.out ./...

latest-release:
	curl -sL https://api.github.com/repos/automoto/devlog/releases/latest | jq -r '.assets[].browser_download_url'
