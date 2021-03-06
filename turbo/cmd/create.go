package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"turbo"
)

var createCmd = &cobra.Command{
	Use:     "create package_path ServiceName",
	Aliases: []string{"c"},
	Short:   "Create a project with runnable HTTP server and gRPC/thrift server",
	Example: "turbo create package/path/to/yourservice YourService -r grpc\n" +
		"'ServiceName' *MUST* be a CamelCase string",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("invalid args")
		}
		// TODO assert that args[1] must be a CamelCase string
		turbo.CreateProject(args[0], args[1], cRpcType)
		return nil
	},
}

var cRpcType string

func init() {
	createCmd.Flags().StringVarP(&cRpcType, "rpctype", "r", "grpc", "[grpc|thrift]")
	RootCmd.AddCommand(createCmd)
}
