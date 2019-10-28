package proto

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"io/ioutil"
	"os"
)

//
//func Example1() {
//	file := "p.proto"
//	imports := []string{"E:\\研发人员项目文件\\Queen-SVN\\Queen-Domain-Layer\\trunk\\src\\queen.com\\m\\stocking\\desk_basic\\pb"}
//	fd, err := ParseFile(file, imports...)
//	if err != nil {
//		// error handling
//	}
//
//	fmt.Printf("%v\n", fd)
//	// Output: {Home:/tmp/fakehome Port:3000 IsProduction:false Inner:{Foo:foobar}}
//}

func Example2() {

	g := generator.New()
	//os.Stdin.WriteString("protoc --go_out=plugins=grpc:. p.proto\n")
	f, err := os.Open("E:\\研发人员项目文件\\Queen-SVN\\Queen-Domain-Layer\\trunk\\src\\queen.com\\m\\stocking\\desk_basic\\pb\\p.proto")
	if err != nil {

	}
	data, err := ioutil.ReadAll(f) //os.Stdin
	g.Request = &plugin.CodeGeneratorRequest{}
	////data = []byte("--go_out=plugins=grpc:. p.proto")
	//data := []byte{}
	d := descriptor.FileDescriptorProto{}
	data = append(data, []byte("\n")...)
	if err := proto.Unmarshal(data, &d); err != nil {
		g.Error(err, "parsing input proto")
	}
	s := "--go_out=plugins=grpc:. p.proto"

	g.Request.Parameter = &s
	g.Request.FileToGenerate = []string{"p.proto"}
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

	// Output: {Home:/tmp/fakehome Port:3000 IsProduction:false Inner:{Foo:foobar}}
}
