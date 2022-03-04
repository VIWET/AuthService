.PHONY: build
build: 
	go build -o build/main cmd/main/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...