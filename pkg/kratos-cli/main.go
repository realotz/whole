package main

import (
	"github.com/realotz/whole/pkg/kratos-cli/internal/proto"
	"github.com/spf13/cobra"
	"log"
)

var (
	version = "v0.0.1"

	rootCmd = &cobra.Command{
		Use:     "kratos-cli",
		Short:   "kratos 开发扩展工具",
		Long:    `Kratos 开发扩展工具`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(proto.CmdProto)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
