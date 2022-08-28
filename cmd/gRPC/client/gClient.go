package main

import (
	profilepb "awesomeNET/proto"
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

	client := profilepb.NewRusProfileServiceClient(conn)

	request := &profilepb.GetInfoRequest{
		Inn: "7716536701",
	}
	
	response, err := client.GetInfo(context.Background(), request)
	if err != nil {
		grpclog.Fatal(err)
	}

	fmt.Println(response)
}
