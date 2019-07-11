package client

import (
	"context"
	pb "github.com/codycollier/grip-go/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
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

func CallEcho(gripcl pb.GripClient, msg string, sleepSeconds int32) {

	// Craft the message
	request := &pb.EchoRequest{Msg: msg, SleepSeconds: sleepSeconds}
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

func CallEchoStream(gripcl pb.GripClient, count int) {

	// Simple echo stream
	stream, err := gripcl.EchoStream(context.Background())
	if err != nil {
		log.Println("Error establishing stream")
		return
	}

	// Wait
	waitc := make(chan struct{})

	// Listener
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("[grip-client] error receiving: %v", err)
			}
			log.Println("[grip-client] recv: ", resp.Msg)
		}
	}()

	for i := 0; i < count; i++ {
		req := &pb.EchoRequest{Msg: strconv.Itoa(i), SleepSeconds: 0}
		stream.Send(req)
		log.Println("[grip-client] send: ", req.Msg)
	}
	<-waitc

}

func CallCompute(gripcl pb.GripClient, computeSeconds int32) {

	// Craft the message
	request := &pb.ComputeRequest{ComputeSeconds: computeSeconds}
	log.Printf("[grip-client] sending: %v", request)

	// Set client context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Make the call
	response, err := gripcl.Compute(ctx, request)
	if err != nil {
		log.Println("[grip-client] Error calling Compute")
	}
	log.Printf("[grip-client] received: %v", response)

}
