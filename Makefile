GO_VERSION := 1.22.0
BIN_DIR := ./bin
BIN_NAME := api

# Setup
.PHONY: install-go init-go
setup: install-go init-go

# TO-DO (rickydodd): Add MacOS support
install-go:
	curl -LO https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz
	sudo tar -xzf go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local
	rm go${GO_VERSION}.linux-amd64.tar.gz

# TO-DO (rickydodd): Add zsh support
init-go:
	echo "export PATH=$$PATH:/usr/local/go/bin" >> $$HOME/.bashrc

# Dev
.PHONY: build tidy
build:
	go build -o ${BIN_DIR}/${BIN_NAME} .

tidy:
	go mod tidy
	gofmt -w .
