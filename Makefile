

proto:
	rm ./examples/*.eva.go
	protoc  -I=./examples --eva_out=plugins=client:./examples hello.proto


grpc:
	protoc -I=./examples --go_out=plugins=grpc:./examples hello.proto

run:
	go install
	make proto