package server

import (
	"context"
	"github.com/codycollier/grip-go/pkg/mockload"
	pb "github.com/codycollier/grip-go/proto"
	"google.golang.org/grpc"
	"log"
	"net"
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
