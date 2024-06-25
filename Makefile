.PHONY: build   

build:
	@go build -o build/userms cmd/main.go 

run: build
	@./bin/userms
