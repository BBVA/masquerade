.PHONY: aceptance test cover clean deps
SHELL := /bin/bash

define gocmd
    @docker run --rm -it -v $$(pwd)/.go:/go -v $$(pwd):/go/src/github.com/BBVA/masquerade golang:1.10.3 go $(1)
endef

all: masquerade

acceptance: masquerade
	$(call gocmd,get github.com/DATA-DOG/godog/cmd/godog)
	$(call gocmd,get github.com/DATA-DOG/godog)
	@docker-compose -f acceptance.yml up -d rabbit s3
	@docker-compose -f acceptance.yml up acceptance s3-test hdfs-test
	@docker-compose -f acceptance.yml stop

masquerade: test
	$(call gocmd,install github.com/BBVA/masquerade/cmd/...)

test: deps
	$(call gocmd,test -race github.com/BBVA/masquerade/pkg/...)

cover: deps
	$(call gocmd,test -race -coverprofile=c.out github.com/BBVA/masquerade/pkg/...)
	$(call gocmd,tool cover -html=c.out -o coverage.html)

deps:
	$(call gocmd,get github.com/ugorji/go/codec)
	$(call gocmd,get github.com/streadway/amqp)
	$(call gocmd,get github.com/spf13/cobra)
