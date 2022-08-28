package main

import (
	pb "awesomeNET/gen"
	"context"
	. "fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const LINK = "https://www.rusprofile.ru/ajax.php?query=7716536701&action=search&cacheKey=%.12f"

type server struct {
	pb.UnimplementedRusProfileServiceServer
}

func (s *server) GetInfo(ctx context.Context, request *pb.GetInfoRequest) (*pb.InfoResponse, error) {
	Println(request.Inn)
	Println(request.GetInn())
	return &pb.InfoResponse{
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
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register
		pb.RegisterRusProfileServiceHandlerServer(context.Background(), mux, &server{})

		// http server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	list, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	Println("Product Svc on", list.Addr().String())

	s := grpc.NewServer()
	pb.RegisterRusProfileServiceServer(s, &server{})
	if err = s.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
