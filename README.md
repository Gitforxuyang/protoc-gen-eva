# protoc-gen-eva
根据proto自动生成代码

仿照go micro的生成代码设置可以将需要的代码生成在一个独立的文件里，分层更合理

使用命令 

protoc  -I=./examples --eva_out=plugins=all:./examples hello.proto

plguins参数可以为server、client、server+client、all 四种，其中server+client等同于all。
输入其它参数则会报错
