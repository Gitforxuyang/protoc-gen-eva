

proto:
	protoc -I=./examples --xmicro_out=plugins=xmicro:./examples/hello hello.proto