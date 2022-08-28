gen:
	protoc --go_out=. --go-grpc_out=. .\proto\rusprofile.proto
clean:
	rm proto/*.go
runC:
	go run cmd/gRPC/client/gClient.go
runS:
	go run cmd/gRPC/server/gServer.go