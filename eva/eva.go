package eva

import (
	"fmt"
	"github.com/Gitforxuyang/protoc-gen-eva/generator"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"strings"
)
const (
	MODE_EVA="eva"
	MODE_SERVER="server"
	MODE_CLIENT="client"
)
type EvaPlugin struct {
	g *generator.Generator
	mode string //模式  eva默认模式 server-只生成server client-只生成client
}

//func (m *EvaPlugin) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName) {
//	m.genImportCode(file)
//}

func init(){
	generator.RegisterPlugin(&EvaPlugin{mode:MODE_EVA})
	generator.RegisterPlugin(&EvaPlugin{mode:MODE_SERVER})
	generator.RegisterPlugin(&EvaPlugin{mode:MODE_CLIENT})
}
func (m *EvaPlugin) Name() string{
	return m.mode
}
func (m *EvaPlugin) Init(g *generator.Generator) {
	m.g=g
}
func (m *EvaPlugin) GenerateImports(file *generator.FileDescriptor, imports map[generator.GoImportPath]generator.GoPackageName){
	if len(file.Service)==0{
		return
	}
	m.genImportCode(file)
}

func (m *EvaPlugin) Generate(files *generator.FileDescriptor){
	if len(files.Service)==0 {
		return
	}
	for _,svc:=range files.Service{
		if m.mode==MODE_EVA{
			m.genServerCode(svc)
			m.genClientCode(svc)
		}
		if m.mode==MODE_SERVER{
			m.genServerCode(svc)
		}
		if m.mode==MODE_CLIENT{
			m.genServerCode(svc)
		}
	}
}

func (m *EvaPlugin) genImportCode(file *generator.FileDescriptor){
	m.g.P("import (")
	m.g.P(`"context"`)
	m.g.P(")")
}

func (m *EvaPlugin) genServerCode(svc *descriptor.ServiceDescriptorProto){
	m.g.P(fmt.Sprintf("type I%sServer interface {",*svc.Name))
	for _,v:=range svc.Method{
		var input string=*v.InputType
		inputArr:=strings.Split(input,".")
		inputArr=inputArr[:copy(inputArr,inputArr[2:])]

		var output string=*v.OutputType
		outputArr:=strings.Split(output,".")
		outputArr=outputArr[:copy(outputArr,outputArr[2:])]
		m.g.P(fmt.Sprintf("%s(ctx context.Context,req *%s) (resp *%s,err error)",
			*v.Name,
			strings.Join(inputArr,"_"),
			strings.Join(outputArr,"_"),
		))
	}
	m.g.P(fmt.Sprintf("}"))
	m.g.P("")
}

func (m *EvaPlugin) genClientCode(svc *descriptor.ServiceDescriptorProto){
	m.g.P(fmt.Sprintf("type I%sClient interface {",*svc.Name))
	for _,v:=range svc.Method{
		var input string=*v.InputType
		inputArr:=strings.Split(input,".")
		inputArr=inputArr[:copy(inputArr,inputArr[2:])]

		var output string=*v.OutputType
		outputArr:=strings.Split(output,".")
		outputArr=outputArr[:copy(outputArr,outputArr[2:])]
		m.g.P(fmt.Sprintf("%s(ctx context.Context,req *%s) (resp *%s,err error)",
			*v.Name,
			strings.Join(inputArr,"_"),
			strings.Join(outputArr,"_"),
		))
	}
	m.g.P(fmt.Sprintf("}"))
}