GO = go

build:
	$(GO) build ./...

test:
	$(GO) test -v -p 1 ./...

tidy:
	$(GO) mod tidy

imports:
	goimports -w .

fmt:
	gofmt -w .

spruce: tidy fmt imports
