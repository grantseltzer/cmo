package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type cmoOptions struct {
	outColor      color.Attribute
	errColor      color.Attribute
	verbose       bool
	combineOutput bool
}

func parseConfigurationOptions(opts *cmoOptions) error {

	// Read in configuration options from the environment
	viper.SetEnvPrefix("CMO")
	viper.SetDefault("STDERR_COLOR", "red")
	viper.SetDefault("STDOUT_COLOR", "green")
	viper.SetDefault("COMBINE_OUPUT", true)
	viper.SetDefault("VERBOSE", false)
	viper.AutomaticEnv()

	specifiedStdoutColor := viper.GetString("STDOUT_COLOR")
	specifiedStderrColor := viper.GetString("STDERR_COLOR")

	opts.combineOutput = viper.GetBool("COMBINE_OUTPUT")
	opts.verbose = viper.GetBool("VERBOSE")

	outColor := colorBindings[strings.ToUpper(specifiedStdoutColor)]
	if outColor == 0 {
		return fmt.Errorf("Unrecognized stdout color: %s. Use 'red', 'green', 'blue', 'yellow', 'cyan', or 'black'",
			specifiedStdoutColor)
	}
	opts.outColor = outColor

	errColor := colorBindings[strings.ToUpper(specifiedStderrColor)]
	if errColor == 0 {
		return fmt.Errorf("Unrecognized stderr color: %s. Use 'red', 'green', 'blue', 'yellow', 'cyan', or 'black'",
			specifiedStderrColor)
	}
	opts.errColor = errColor

	return nil
}

var colorBindings = map[string]color.Attribute{
	"RED":    color.FgRed,
	"BLUE":   color.FgBlue,
	"GREEN":  color.FgGreen,
	"YELLOW": color.FgYellow,
	"CYAN":   color.FgCyan,
	"BLACK":  color.FgBlack,
}
