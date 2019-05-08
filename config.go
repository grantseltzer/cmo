package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var colorBindings = map[string]color.Attribute{
	"RED":    color.FgRed,
	"BLUE":   color.FgBlue,
	"GREEN":  color.FgGreen,
	"YELLOW": color.FgYellow,
	"CYAN":   color.FgCyan,
	"BLACK":  color.FgBlack,
}

func parseConfigurationOptions(opts *cmoOptions) error {
	viper.SetEnvPrefix("CMO")
	viper.SetDefault("STDERR_COLOR", "red")
	viper.SetDefault("STDOUT_COLOR", "green")
	viper.SetDefault("COMBINE_OUPUT", true)
	viper.AutomaticEnv()

	outColor := colorBindings[strings.ToUpper(viper.GetString("STDOUT_COLOR"))]
	if outColor == 0 {
		return fmt.Errorf("Unrecognized stdout color: %s. Use 'red', 'green', 'blue', 'yellow', 'cyan', or 'black'",
			viper.GetString("STDOUT_COLOR"))
	}
	opts.outColor = outColor

	errColor := colorBindings[strings.ToUpper(viper.GetString("STDERR_COLOR"))]
	if errColor == 0 {
		return fmt.Errorf("Unrecognized stderr color: %s. Use 'red', 'green', 'blue', 'yellow', 'cyan', or 'black'",
			viper.GetString("STDERR_COLOR"))
	}
	opts.errColor = errColor

	return nil
}
