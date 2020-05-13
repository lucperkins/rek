GO = go

build:
	$(GO) build ./...

test:
	$(GO) test ./...

tidy:
	$(GO) mod tidy

imports:
	goimports -w .

fmt:
	gofmt -w .

spruce: tidy fmt imports
