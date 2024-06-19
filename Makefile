.PHONY: build   

build:
	@go build -o bin/authsrv cmd/main.go 

run: build
	@./bin/authsrv
