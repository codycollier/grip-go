package server

import (
	"context"
	"fmt"
	pb "github.com/codycollier/grip-go/proto"
	// "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"net"
)

type gripServer struct {
}

func newGripServer() *gripServer {
	return &gripServer{}
}

func (s *gripServer) Echo(ctx context.Context, echo *pb.EchoRequest) (*pb.EchoResponse, error) {
	resp := &pb.EchoResponse{Msg: "pong"}
	return resp, nil
}

func (s *gripServer) Compute(ctx context.Context, creq *pb.ComputeRequest) (*pb.ComputeResponse, error) {
	resp := &pb.ComputeResponse{Msg: "123abc"}
	return resp, nil
}

func StartGripServer() {

	fmt.Println("[gripd] Hello world!")

	// Setup the gRPC server and start
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("[gripd] Error listening on 8080")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGripServer(grpcServer, newGripServer())
	grpcServer.Serve(listener)

}
