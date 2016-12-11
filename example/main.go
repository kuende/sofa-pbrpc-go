package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/kuende/sofa-pbrpc-go/sofa"
	"github.com/kuende/sofa-pbrpc-go/sofa/discovery"

	echo "github.com/kuende/sofa-pbrpc-go/example/generated/sofa_pbrpc_test"
)

var (
	serverAddr string
	message    string
)

func sofaRPC() {
	clientConn, err := sofa.NewClient(sofa.ClientOpts{
		TransportType:     sofa.TCPTransportType,
		DiscoveryProvider: discovery.NewStaticProvider([]string{serverAddr}),
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
		os.Exit(1)
	}

	client := echo.NewEchoServerClient(clientConn)

	response, err := client.Echo(context.Background(), &echo.EchoRequest{Message: proto.String(message)})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to send rpc: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Received response for sofa-pbrpc: %v\n", response)
}

func httpRPC() {
	clientConn, err := sofa.NewClient(sofa.ClientOpts{
		TransportType:     sofa.HTTPTransportType,
		DiscoveryProvider: discovery.NewStaticProvider([]string{serverAddr}),
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
		os.Exit(1)
	}
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
