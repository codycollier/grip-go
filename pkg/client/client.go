package client

import (
	"context"
	pb "github.com/codycollier/grip-go/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetNewGripClient(addr string) (pb.GripClient, *grpc.ClientConn) {

	// Setup the gRPC connection & grip client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("[grip-client] Error dialing server")
	}
	gripcl := pb.NewGripClient(conn)

	// Return the conn too, so defer conn.Close() can be called
	return gripcl, conn
}

func CallEcho(gripcl pb.GripClient, msg string, sleep int) {

	// Craft the message
	request := &pb.EchoRequest{Msg: "foo", SleepSeconds: 0}
	log.Printf("[grip-client] sending: %v", request)

	// Set client context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Make the call
	response, err := gripcl.Echo(ctx, request)
	if err != nil {
		log.Println("[grip-client] Error calling Echo")
	}
	log.Printf("[grip-client] received: %v", response)

}
