package proto

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"io/ioutil"
)

// https://github.com/tallstoat/pbparser

func loadFileDesc(filename string) (*descriptor.FileDescriptorProto, []byte) {
	enc := proto.FileDescriptor(filename)
	if enc == nil {
		panic(fmt.Sprintf("failed to find fd for file: %v", filename))
	}
	fd, err := decodeFileDesc(enc)
	if err != nil {
		panic(fmt.Sprintf("failed to decode enc: %v", err))
	}
	b, err := proto.Marshal(fd)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal fd: %v", err))
	}
	return fd, b
}
func decodeFileDesc(enc []byte) (*descriptor.FileDescriptorProto, error) {
	raw, err := decompress(enc)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress enc: %v", err)
	}

	fd := new(descriptor.FileDescriptorProto)
	if err := proto.Unmarshal(raw, fd); err != nil {
		return nil, fmt.Errorf("bad descriptor: %v", err)
	}
	return fd, nil
}

func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("bad gzipped descriptor: %v", err)
	}
	return out, nil
}

//func test() {
//	g := generator.New()
//
//	data, err := ioutil.ReadAll(os.Stdin)
//	if err != nil {
//		g.Error(err, "reading input")
//	}
//
//	if err := proto.Unmarshal(data, g.Request); err != nil {
//		g.Error(err, "parsing input proto")
//	}
//
//	generator.Annotate()
//}
//func ParseRequest(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
//	input, err := ioutil.ReadAll(r)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read code generator request: %v", err)
//	}
//	req := new(plugin.CodeGeneratorRequest)
//	if err = proto.Unmarshal(input, req); err != nil {
//		return nil, fmt.Errorf("failed to unmarshal code generator request: %v", err)
//	}
//	return req, nil
//}
//
//func ParseFile(filename string, paths ...string) (*descriptor.FileDescriptorSet, error) {
//	return parseFile(filename, false, true, paths...)
//}
// https://blog.csdn.net/lufeng20/article/details/8736584
//func parseFile(filename string, includeSourceInfo bool, includeImports bool, paths ...string) (*descriptor.FileDescriptorSet, error) {
//	args := []string{"--proto_path=" + strings.Join(paths, ":")}
//	if includeSourceInfo {
//		args = append(args, "--include_source_info")
//	}
//	if includeImports {
//		args = append(args, "--include_imports")
//	}
//	args = append(args, "--descriptor_set_out=/dev/stdout")
//	args = append(args, filename)
//	cmd := exec.Command("protoc", args...)
//	cmd.Env = []string{}
//	data, err := cmd.CombinedOutput()
//	if err != nil {
//		return nil, err
//	}
//	fileDesc := &descriptor.FileDescriptorSet{}
//	if err := proto.Unmarshal(data, fileDesc); err != nil {
//		return nil, err
//	}
//	return fileDesc, nil
//}
