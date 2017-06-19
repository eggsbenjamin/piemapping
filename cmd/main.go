package main

import (
	"github.com/eggsbenjamin/piemapping/commons"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
)

var logr commons.LevelledLogWriter

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("piemapping")
	logs := viper.GetBool("logging")
	if logs == true {
		lf := viper.GetString("log_format")
		df := commons.GetLogLevelId(lf)
		logr = commons.NewLogger("Piemapping", df)
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
