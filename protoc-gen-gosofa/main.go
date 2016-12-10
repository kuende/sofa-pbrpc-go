package main

import (
	"github.com/gogo/protobuf/vanity/command"
	_ "github.com/kuende/sofa-pbrpc-go/protoc-gen-gosofa/sofa"
)

func main() {
	command.Write(command.Generate(command.Read()))
}
