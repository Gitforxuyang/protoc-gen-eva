package xmicro

import (
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type XMicroPlugin struct {
	g *generator.Generator
}

//func (m *XMicroPlugin) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
//	m.genImportCode(file)
//}

func init(){
	generator.RegisterPlugin(new(XMicroPlugin))
}
func (m *XMicroPlugin) Name() string{
	return "eva"
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
	m.g.P(fmt.Sprintf("// %s",*file.Name))
}
func (m *XMicroPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto){
	m.g.P("// todo :genServiceCode ")
	m.g.P(fmt.Sprintf("// %s",*svc.Name))
}