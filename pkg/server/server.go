package server

import (
	"context"
	pb "github.com/codycollier/grip-go/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type gripServer struct {
}

func (s *gripServer) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("[gripd] received: %v", req)
	resp := &pb.EchoResponse{Msg: req.Msg}
	log.Printf("[gripd] sending: %v", resp)
	return resp, nil
}

func (s *gripServer) Compute(ctx context.Context, req *pb.ComputeRequest) (*pb.ComputeResponse, error) {
	log.Printf("[gripd] received: %v", req)
	resp := &pb.ComputeResponse{Msg: "123abc"}
	log.Printf("[gripd] sending: %v", resp)
	return resp, nil
}

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
