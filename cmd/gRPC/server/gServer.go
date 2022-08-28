package main

import (
	profilepb "awesomeNET/proto"
	"context"
	. "fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const LINK = "https://www.rusprofile.ru/ajax.php?query=7716536701&action=search&cacheKey=%.12f"

type server struct {
	profilepb.UnimplementedRusProfileServiceServer
}

func (s *server) GetInfo(ctx context.Context, request *profilepb.GetInfoRequest) (*profilepb.InfoResponse, error) {
	Println(request.Inn)
	Println(request.GetInn())
	return &profilepb.InfoResponse{
		Inn:     request.Inn,
		Kpp:     request.GetInn(),
		Name:    "OOO MYCOMPANI22222",
		CeoName: "Verteletsky Roman",
	}, nil
}

func main() {
	startGRPC()
}

func startGRPC() {
	list, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	Println("Product Svc on", list.Addr().String())

	s := grpc.NewServer()
	profilepb.RegisterRusProfileServiceServer(s, &server{})
	if err = s.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
