.PHONY:
	build move-to-bin install

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
	golint

test:
	go test