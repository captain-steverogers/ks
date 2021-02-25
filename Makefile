HELL=/bin/bash
BINARY_NAME:="ks"
GOPATH:="${HOME}/go"

.PHONY: install
install: ## Install the binary 
	@go build -o ${GOPATH}/bin/${BINARY_NAME}