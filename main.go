package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {

	var opts cmoOptions
	err := parseConfigurationOptions(&opts)
	if err != nil {
		if opts.verbose {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}

	rootCmd := &cobra.Command{
		Use:                "cmo [command and args]",
		Short:              "color coordinates stdout/stderr of the command",
		DisableFlagParsing: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(args)
		},
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		rootCmd.Help()
		os.Exit(0)
	}

	err = rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}

func (c cmoOptions) run(args []string) error {

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = newDirectWriter(os.Stdout, c.outColor)

	// whether or not to print stderr to stdout
	// to make pipes worth. Still color coded.
	if c.combineOutput {
		cmd.Stderr = newDirectWriter(os.Stdout, c.errColor)
	} else {
		cmd.Stderr = newDirectWriter(os.Stderr, c.errColor)
	}

	return cmd.Run()
}
