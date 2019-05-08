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
	cmd.Stdout = newDirectWriter(c.outColor)
	cmd.Stderr = newDirectWriter(c.errColor)

	//TODO:
	// Add ability to "combine output", meaning stderr would also
	// be printed to stdout except with its own color. Need to have
	// a multi-writer with which multiple writers can call their
	// own write methods which create a formatted (colored) string
	// and write to an underlying writer for stdout
	//

	return cmd.Run()
}
