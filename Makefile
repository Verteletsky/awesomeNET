create:
	protoc --proto_path=. proto/*.proto --go_out=gen/
	protoc --proto_path=. proto/*.proto --go-grpc_out=gen/
	protoc -I . --grpc-gateway_out ./gen/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true .\proto\rusprofile.proto
clean:
	rm gen/proto/*.go
runC:
	go run cmd/gRPC/client/gClient.go
runS:
	go run cmd/gRPC/server/gServer.go
tidy:
	go mod tidy