package main

import (
	"fmt"
	helloworld "github.com/openjw/genter/x/protobuf/v2"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	p1 := helloworld.HelloReply{Message: "123"}
	out, err := proto.Marshal(&p1)
	p2 := helloworld.HelloReply{}
	if err := proto.Unmarshal(out, &p2); err != nil {

	}
	fmt.Println(p2.Message)

	gen, err := protogen.Options{}.New(&pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{
			{
				Name:    proto.String("helloworld.proto"),
				Syntax:  proto.String(protoreflect.Proto3.String()),
				Package: proto.String("helloworld"),
			},
		},
	})
	if err != nil {

	}

	for i, f := range gen.Files {
		if got, want := string(f.GoPackageName), "goproto_testdata"; got != want {
			fmt.Errorf("gen.Files[%d].GoPackageName = %v, want %v", i, got, want)
		}
		if got, want := string(f.GoImportPath), "testdata/go_package"; got != want {
			fmt.Errorf("gen.Files[%d].GoImportPath = %v, want %v", i, got, want)
		}
	}

	//g := gen.NewGeneratedFile("foo.go", "golang.org/x/foo")

}
