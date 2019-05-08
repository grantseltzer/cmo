package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type cmoOptions struct {
	outColor      color.Attribute
	errColor      color.Attribute
	combineOutput bool
}

func main() {

	var opts cmoOptions
	err := parseConfigurationOptions(&opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error()) //FIXME: put behind verbose flag
	}

	rootCmd := &cobra.Command{
		Use:                "cmo [command and args]",
		Short:              "color coordinates stdout/stderr of the command",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(args)
		},
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func (c cmoOptions) run(args []string) error {

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = newDirectWriter(c.outColor)
	cmd.Stderr = newDirectWriter(c.errColor)

	return cmd.Run()
}
