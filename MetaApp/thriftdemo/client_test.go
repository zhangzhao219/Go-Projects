package main

import (
	"context"
	"fmt"
	"sample/gen-go/Sample"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
)

var ctx = context.Background()

func GetClient() *Sample.GreeterClient {
	addr := ":9090"
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
	}

	//protocol
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	//no buffered
	transportFactory := thrift.NewTTransportFactory()

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("error running client:", err)
	}

	if err := transport.Open(); err != nil {
		fmt.Println("error running client:", err)
	}

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	client := Sample.NewGreeterClient(thrift.NewTStandardClient(iprot, oprot))
	return client
}

// GetUser
func TestGetUser(t *testing.T) {
	client := GetClient()
	rep, err := client.GetUser(ctx, 100)
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}

// SayHello
func TestSayHello(t *testing.T) {
	client := GetClient()
	rep, err := client.SayHello(ctx, &Sample.User{
		Name:    "thrift",
		Address: "address",
	})
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}
