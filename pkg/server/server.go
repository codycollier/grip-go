package server

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/codycollier/grip-go/pkg/mockload"
	pb "github.com/codycollier/grip-go/proto"
	"google.golang.org/grpc"
)

// Core struct representing the gRPC Grip service
type gripServer struct{}

// Main function for Grip.Echo() endpoint
func (s *gripServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("[gripd] received: %v", req)

	resp := &pb.EchoResponse{Msg: req.Msg}

	log.Printf("[gripd] sending: %v", resp)
	return resp, nil
}

// Main function for Grip.EchoStream() endpoint
func (s *gripServer) EchoStream(stream pb.Grip_EchoStreamServer) error {

	// Iterate over stream of EchoRequest messages
	for {
		req, err := stream.Recv()

		// End of stream
		if err == io.EOF {
			break
		}

		// Error
		if err != nil {
			return err
		}

		// Echo the string back
		log.Println("[gripd] recv: ", req.Msg)
		resp := &pb.EchoResponse{Msg: req.Msg}
		if err := stream.Send(resp); err != nil {
			return err
		}
		log.Println("[gripd] send: ", resp.Msg)

	}

	// End of response stream
	return nil
}

// Main function for Grip.Compute() endpoint
func (s *gripServer) Compute(ctx context.Context, req *pb.ComputeRequest) (*pb.ComputeResponse, error) {
	log.Printf("[gripd] received: %v", req)

	computeMsg := mockload.ComputeLoad(req.ComputeSeconds)
	resp := &pb.ComputeResponse{Msg: computeMsg}

	log.Printf("[gripd] sending: %v", resp)
	return resp, nil
}

// Start the gRPC server & register the defined Grip service
func StartGripServer(addr string) {

	// Setup the gRPC server and start
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("[gripd] Error listening on %s", addr)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGripServer(grpcServer, &gripServer{})

	// Listen forever
	log.Printf("[gripd] Listening on %v", addr)
	grpcServer.Serve(listener)

}
