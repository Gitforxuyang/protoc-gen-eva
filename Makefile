

proto:
	protoc  -I=./examples --eva_out=plugins=eva+grpc:./examples/hello hello.proto


grpc:
	protoc -I=./examples --go_out=plugins=grpc:./examples/hello hello.proto