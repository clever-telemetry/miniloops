package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/clever-telemetry/miniloops/pkg/cli/cmd"
)

func init() {
	cobra.OnInitialize(cobraInit)
}

func cobraInit() {
	// Defaults
	viper.SetDefault("kubernetes.client.qps", 10)
	viper.SetDefault("kubernetes.client.burst", 20)
	viper.SetDefault("manager.metrics.address", "127.0.0.1:9100")
	viper.SetDefault("manager.address", "0.0.0.0")
	viper.SetDefault("manager.port", 8080)
}

func Execute() error {
	return cmd.RootCmd.Execute()
}
