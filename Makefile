

proto:
	protoc  -I=./examples --eva_out=plugins=eva:./examples hello.proto


grpc:
	protoc -I=./examples --go_out=plugins=grpc:./examples/hello hello.proto