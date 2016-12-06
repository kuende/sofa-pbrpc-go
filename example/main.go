package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/kuende/sofa-pbrpc-go/sofa"

	echo "github.com/kuende/sofa-pbrpc-go/example/generated/sofa_pbrpc_test"
)

var (
	serverAddr string
	message    string
)

func sofaRPC() {
	transport, err := sofa.NewTCPTransport([]string{serverAddr}, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize transport: %v\n", err)
		os.Exit(1)
	}

	clientConn := sofa.NewClient(transport)
	client := echo.NewEchoServerClient(clientConn)

	response, err := client.Echo(context.Background(), &echo.EchoRequest{Message: proto.String(message)})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send rpc: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Received response for sofa-pbrpc: %v\n", response)
}

func httpRPC() {
	transport, _ := sofa.NewHTTPTransport(serverAddr)
	clientConn := sofa.NewClient(transport)
	client := echo.NewEchoServerClient(clientConn)

	response, err := client.Echo(context.Background(), &echo.EchoRequest{Message: proto.String(message)})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send rpc: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Received response for http: %v\n", response)
}

func main() {
	flag.StringVar(&serverAddr, "s", "localhost:12321", "server address")
	flag.StringVar(&message, "m", "Hello from qinzuoyan01", "message to send")
	flag.Parse()

	sofaRPC()
	httpRPC()
}
