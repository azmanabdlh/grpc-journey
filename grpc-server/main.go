package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "golang-grpc-example/grpc-proto/helloworld"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Hello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: "Hello world",
	}, nil
}

func (s *server) HelloWithStream(stream pb.HelloWorld_HelloWithStreamServer) error {

	for i := 1; i <= 10; i++ {
		err := stream.Send(&pb.Response{
			Message: fmt.Sprintf("Data %d", i),
		})
		if err != nil {
			fmt.Println("Error Send:", err)
			break
		}

		time.Sleep(time.Second * 3)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50059")
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}
}
