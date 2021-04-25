package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
	"strings"
)

const version = "0.0.1"

func main() {
	versionF := flag.Bool("version", false, "version")
	flag.Parse()
	if *versionF {
		fmt.Printf("protoc-gen-openapi %v\n", version)
		return
	}
	var flags flag.FlagSet
	flags.String("path", "", "path")
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		param := make(map[string]string)
		for _, p := range strings.Split(gen.Request.GetParameter(), ",") {
			if i := strings.Index(p, "="); i < 0 {
				param[p] = ""
			} else {
				param[p[0:i]] = p[i+1:]
			}
		}
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			generateFile(gen, f, param)
		}
		return nil
	})
}
