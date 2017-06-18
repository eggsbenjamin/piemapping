package main

import (
	"github.com/eggsbenjamin/piemapping/commons"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logr commons.LevelledLogWriter

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("piemapping")
	logs := viper.GetBool("logging")
	if logs == true {
		logr = commons.NewLogger("Piemapping", 1)
		return
	}
	logr = &commons.NoopLogger{}
}

func main() {
	cmd := &cobra.Command{Use: "piemapping"}
	cmd.AddCommand(migrate())
	cmd.AddCommand(run())
	cmd.Execute()
}
