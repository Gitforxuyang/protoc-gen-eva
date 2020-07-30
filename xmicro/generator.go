package xmicro

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type XMicroPlugin struct {
	g *generator.Generator
}

func init(){
	generator.RegisterPlugin(new(XMicroPlugin))
}
func (m *XMicroPlugin) Name() string{
	return "xmicro"
}
func (m *XMicroPlugin) Init(g *generator.Generator) {
	m.g=g
}
func (m *XMicroPlugin) GenerateImports(files *generator.FileDescriptor){
	if len(files.Service)>0{
		m.genImportCode(files)
	}
}

func (m *XMicroPlugin) Generate(files *generator.FileDescriptor){
	for _,svc:=range files.Service{
		m.genServiceCode(svc)
	}
}

func (m *XMicroPlugin) genImportCode(file *generator.FileDescriptor){
	m.g.P("// todo: genImportCode")
}
func (m *XMicroPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto){
	m.g.P("// todo :genServiceCode ")
}