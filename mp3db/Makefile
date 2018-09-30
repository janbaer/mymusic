install:
	@go get -d -u github.com/golang/dep
	@dep ensure
.PHONY: install

run:
	@go run main.go serve
.PHONY: run

test:
	@go test -v -cover ./...
.PHONY: test

