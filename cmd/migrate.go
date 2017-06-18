package main

import (
	"io/ioutil"

	"github.com/eggsbenjamin/piemapping/repository"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//	base migrate command
func migrate() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate",
		Run: func(cmd *cobra.Command, args []string) {
			logr.Info("Migrating...")
		},
	}
	cmd.AddCommand(up())
	return
}

//	'up' command for creating/seeding db from file
func up() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Up",
		Run: func(cmd *cobra.Command, args []string) {
			var (
				raw  []byte
				err  error
				path = viper.GetString("db_migrations_path")
			)
			if raw, err = ioutil.ReadFile(path); err != nil {
				panic(err.Error())
			}
			logr.Infof("Executing script '%s'...", path)
			var (
				params = map[string]string{"multiStatements": "true"}
				conn   = repository.NewConnection(logr, params)
				sql    = string(raw)
			)
			if _, err = conn.Query(sql); err != nil {
				panic(err.Error())
			}
			logr.Info("Migrations Complete.")
		},
	}
}
