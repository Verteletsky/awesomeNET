package main

import (
	pb "awesomeNET/gen"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRusProfileServiceClient(conn)

	request := &pb.GetInfoRequest{
		Inn: "7716536701",
	}

	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		grpclog.Fatal(err)
	}

	fmt.Println(response)
}
