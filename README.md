# sofa-pbrpc-go

This package implements basic Golang client for [sofa-pbrpc](https://github.com/baidu/sofa-pbrpc)

*This is a very alpha version, expect APIs to break without any notice*

### Example:

See example directory for a complete example implementing Go client for [example echo server](https://github.com/baidu/sofa-pbrpc/tree/master/sample/echo).

### Usage

Install sofa-pbrpc-go protoc plugin:

```bash
$ go get -u github.com/kuende/sofa-pbrpc-go/protoc-gen-gosofa
```

Generate Go files from protobuf IDL:

```bash
$ protoc -I . *.proto --gosofa_out=plugins=sofa:generated/sofa_pbrpc_test
```

### TODO

- [x] basic request/response with timeout
- [X] HTTP/JSON transport
- [x] code generation tool
- [x] SeedProvider for TCP / HTTP transports
- [ ] return error codes to caller
- [ ] load balancing
- [ ] fault tolerance + backoff
- [ ] implement better timeout control
- [ ] implement compression
- [ ] implement test mocking interface
- [ ] RPC server
- [ ] tests, tests, tests
