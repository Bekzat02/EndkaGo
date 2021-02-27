genG:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculatorpb/*.go
clean:
	rm calculatorpb/*.go
runs:
	go run go build calculator_server/server.go
runc:
	go run go build calculator_client/client.go
