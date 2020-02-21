.PHONY:
	build move-to-bin install

build:
	@echo "Building the binary..."
	@go build -o devlog main.go output.go util.go
	@echo "Done."

move-to-bin:
	@echo "Moving binary to /usr/local/bin"
	@sudo cp devlog /usr/local/bin
	@echo "Done."

install:
	make build && make move-to-bin

test:
	go test