package main

import (
	_ "github.com/Gitforxuyang/protoc-gen-eva/eva"
	"github.com/golang/protobuf/proto"
	"github.com/Gitforxuyang/protoc-gen-eva/generator"
	"io/ioutil"
	"os"
)

func main(){
	g:=generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
	g.BuildTypeNameMap()

	g.GenerateAllFiles()

	//g.Fail(*g.Response.File[0].Name)
	//name:="examples/hello.micro.go"
	//// Send back the results.
	//g.Response.GetFile()[0].Name=&name
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}

