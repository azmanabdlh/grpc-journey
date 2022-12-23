package main

import (
	"context"
	"fmt"

	pb "golang-grpc-example/grpc-proto/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50059", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return
	}
	defer conn.Close()

	client := pb.NewHelloWorldClient(conn)
	ctx := context.Background()

	// in := pb.Request{}
	// resp, err := client.Hello(ctx, &in)
	// if err != nil {
	// 	fmt.Println("Error response:", err)
	// 	return
	// }

	// fmt.Printf("Response %+v\n", resp.Message)

	stream, err := client.HelloWithStream(ctx)
	if err != nil {
		fmt.Println("Error response stream:", err)
		return
	}

	for i := 0; i < 5; i++ {
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println("Error response:", err)
			break
		}

		fmt.Printf("Response %+v\n", resp.GetMessage())
	}

	stream.CloseSend()
}
