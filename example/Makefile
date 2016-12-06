.PHONY: proto build

all: build

proto:
	protoc -I . *.proto --gogo_out=generated/sofa_pbrpc_test

build:
	mkdir -p bin
	go build -o bin/echo-client
