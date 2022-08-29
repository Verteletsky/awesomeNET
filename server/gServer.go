package main

import (
	pb "awesomeNET/gen"
	"awesomeNET/models"
	"context"
	"encoding/json"
	"errors"
	. "fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"strings"
	"unicode/utf8"
)

const LINK = "https://www.rusprofile.ru/ajax.php?query=%s&action=search"

type server struct {
	pb.UnimplementedRusProfileServiceServer
}

func (s *server) GetInfo(ctx context.Context, request *pb.GetInfoRequest) (*pb.InfoResponse, error) {
	countInString := utf8.RuneCountInString(request.GetInn())
	if countInString != 12 && countInString != 10 {
		return nil, errors.New("12 цифр ИП, 10 цифр ООО")
	}
	get, err := http.Get(Sprintf(LINK, request.Inn))
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()

	v := &models.Profiles{}
	err = json.NewDecoder(get.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	if v.Success == false {
		if v.Code == 255 {
			return nil, errors.New(v.Message)
		} else {
			return nil, errors.New(v.Message)
		}
	}

	if v.Count == 0 {
		return nil, errors.New("компания не найдена")
	}

	v.Profile[0].Inn = strings.Trim(v.Profile[0].Inn, "~!")

	return &pb.InfoResponse{
		Inn:     v.Profile[0].Inn,
		Kpp:     v.Profile[0].Kpp,
		Name:    v.Profile[0].Name,
		CeoName: v.Profile[0].CeoName,
	}, nil
}

func main() {
	startGRPC()
}

func startGRPC() {
	go func() {
		// serveMux
		serveMux := runtime.NewServeMux()

		// register
		pb.RegisterRusProfileServiceHandlerServer(context.Background(), serveMux, &server{})

		muxHttp := http.NewServeMux()
		muxHttp.Handle("/get/", serveMux)

		// http server
		log.Fatalln(http.ListenAndServe("localhost:8081", muxHttp))
	}()

	list, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	Println("Start gServer on ", list.Addr().String())

	s := grpc.NewServer()
	pb.RegisterRusProfileServiceServer(s, &server{})
	if err = s.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
