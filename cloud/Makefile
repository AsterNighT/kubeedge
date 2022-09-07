DOCKER_REGISTRY ?= "10.10.101.109:5000"

IMAGE_TAG ?= "apiserver"

.PHONY: default all image binary push

default: all

all: binary image push

binary: bin/apiserver

bin/apiserver: cmd/apiserver/main.go
	go build -o bin/apiserver cmd/apiserver/main.go

image: binary
	docker build . -t ${DOCKER_REGISTRY}/${IMAGE_TAG}

push: image
	docker push ${DOCKER_REGISTRY}/${IMAGE_TAG}