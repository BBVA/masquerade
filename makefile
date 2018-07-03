.PHONY: aceptance test cover clean deps
SHELL := /bin/bash

all: masquerade

acceptance: masquerade
	@go get github.com/DATA-DOG/godog/cmd/godog
	@go get github.com/DATA-DOG/godog
	@cd internal/features && godog -t "~@wip" .
	@#cd internal/features && godog -t "@dev" .

masquerade: test
	@go install ./cmd/...

test: clean
	@go get github.com/ugorji/go/codec
	@go test -race ./pkg/...

cover:
	@go test -race -coverprofile=c.out ./pkg/... && go tool cover -html=c.out -o coverage.html

clean:
	@rm c.out || echo "but its ok"