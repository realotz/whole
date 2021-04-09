package server

import (
	"fmt"
	"github.com/realotz/whole/pkg/kratos-cli/internal/base"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// CmdServer represents the source command.
	CmdServer = &cobra.Command{
		Use:   "server",
		Short: "Generate the proto server code",
		Long:  "Generate the proto server code. Example: kratos proto server helloworld.proto",
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please enter the proto file or directory")
		return
	}
	var (
		err   error
		proto = strings.TrimSpace(args[0])
	)
	if err = look("protoc-gen-go", "protoc-gen-go-grpc", "protoc-gen-go-http", "protoc-gen-go-errors"); err != nil {
		// update the kratos plugins
		cmd := exec.Command("kratos", "upgrade")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}
	if strings.HasSuffix(proto, ".proto") {
		err = generate(proto, args)
	} else {
		err = walk(proto, args)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func look(name ...string) error {
	for _, n := range name {
		if _, err := exec.LookPath(n); err != nil {
			return err
		}
	}
	return nil
}

func walk(dir string, args []string) error {
	if dir == "" {
		dir = "."
	}
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if ext := filepath.Ext(path); ext != ".proto" {
			return nil
		}
		return generate(path, args)
	})
}

// generate is used to execute the generate command for the specified proto file
func generate(proto string, localArgs []string) error {
	args := []string{
		"--proto_path=.",
		"--proto_path=" + filepath.Join(base.KratosMod(), "api"),
		"--proto_path=" + filepath.Join(base.KratosMod(), "third_party"),
		"--proto_path=" + filepath.Join(os.Getenv("GOPATH"), "src"),
		"--kratos-server_out=path=" + localArgs[1] + ",paths=source_relative:" + localArgs[1],
	}
	_, err := os.Stat("vendor")
	if err == nil {
		args = append(args, "--proto_path=./vendor")
	}
	args = append(args, proto)
	fd := exec.Command("protoc", args...)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	//fd.Dir = path
	if err := fd.Run(); err != nil {
		fmt.Printf("comand: protoc %s \n", strings.Join(args, " "))
		return err
	}
	fmt.Printf("proto: %s\n", proto)
	fmt.Printf("comand: protoc %s \n", strings.Join(args, " "))
	return nil
}
