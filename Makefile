

proto:
	-rm ./examples/*.eva.go
	protoc -I=./examples/googleapis   -I=./examples --eva_out=plugins=all:./examples hello.proto


grpc:
	protoc  -I=./examples/googleapis -I=./examples --go_out=plugins=grpc:./examples hello.proto

run:
	go install
	make proto