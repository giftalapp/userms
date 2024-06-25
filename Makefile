.PHONY: build   

build:
	@go build -o build/authsrv cmd/main.go 

run: build
	@./bin/userms
