package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

type cmoOptions struct {
	outColor      string
	errColor      string
	combineOutput bool
}

func main() {

	var opts cmoOptions

	rootCmd := &cobra.Command{
		Use:   "cmo [command and args]",
		Short: "color coordinates stdout/stderr of the command",
		Run: func(cmd *cobra.Command, args []string) {

			viper.SetEnvPrefix("CMO")
			viper.SetDefault("STDERR_COLOR", "red")
			viper.SetDefault("STDOUT_COLOR", "green")
			viper.SetDefault("COMBINE_OUPUT", true)
			viper.AutomaticEnv()

			opts.run(args)
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func (c cmoOptions) run(args []string) {

}
