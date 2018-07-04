.PHONY: aceptance test cover clean deps
SHELL := /bin/bash

define gocmd
    @docker run --rm -it -v $$(pwd)/.go:/go -v $$(pwd):/go/src/github.com/BBVA/masquerade golang:1.10.3 go $(1)
endef

all: masquerade

acceptance: masquerade
	$(call gocmd,get github.com/DATA-DOG/godog/cmd/godog)
	$(call gocmd,get github.com/DATA-DOG/godog)
	docker rm -f rabbit || echo but its ok
	docker network rm masqnet || echo but its ok
	docker network create masqnet
	docker run -d --name rabbit --net masqnet --hostname rabbit \
		-e RABBITMQ_DEFAULT_USER=guest -e RABBITMQ_DEFAULT_PASS=guest rabbitmq:3.7.6
	sleep 10
	docker run --rm -it --net masqnet \
		-v $$(pwd)/.go:/go -v $$(pwd):/go/src/github.com/BBVA/masquerade -w /go/src/github.com/BBVA/masquerade/internal/features golang:1.10.3 godog -t "~@wip" . 
	docker rm -f rabbit || echo but its ok
	docker network rm masqnet || echo but its ok

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
