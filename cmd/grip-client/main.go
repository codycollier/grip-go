package main

import (
	"context"
	"log"
	// "github.com/codycollier/grip-go/pkg/client"
	pb "github.com/codycollier/grip-go/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {

	log.Println("[grip-client] Hello world!")

	// gripcl := client.NewGripClient("localhost:8080")

	// Setup the gRPC connection
	var addr = "127.0.0.1:8080"
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("[grip-client] Error dialing server")
	}
	defer conn.Close()

	// Setup the grip client and common items
	gripcl := pb.NewGripClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Make a call to Echo()
	request := &pb.EchoRequest{Msg: "foo", SleepSeconds: 0}
	log.Printf("[grip-client] sending: %v", request)
	response, err := gripcl.Echo(ctx, request)
	if err != nil {
		log.Println("[grip-client] Error calling Echo")
	}
	log.Printf("[grip-client] received: %v", response)

}
