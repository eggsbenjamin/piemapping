package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	env string
)

func init() {
	//	config file
	viper.AddConfigPath("../config")
	viper.SetConfigName("app")
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	//	env vars
	viper.AutomaticEnv()
	prefix := viper.GetString("common.name")
	viper.SetEnvPrefix(prefix)
	env = viper.GetString("env")
	//	logging
	log.SetPrefix("[Info]")
}

//	assemble components and construct the db connection string
func getConnectionString() string {
	DBHost := viper.GetString(fmt.Sprintf("%s.db.host", env))
	DBPort := viper.GetString(fmt.Sprintf("%s.db.port", env))
	DBName := viper.GetString(fmt.Sprintf("%s.db.name", env))
	DBUser := viper.GetString("db_user")
	DBPwd := viper.GetString("db_password")
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		DBUser,
		DBPwd,
		DBHost,
		DBPort,
		DBName,
	)
}

//	connects to the db and returns the connection object.
//	this function panics rather than errors due it's criticality.
func NewConnection() (DB *sql.DB) {
	var (
		err     error
		DBType  = viper.GetString(fmt.Sprintf("%s.db.type", env))
		connStr = getConnectionString()
	)
	log.Println("Connecting to DB...")
	if DB, err = sql.Open(DBType, connStr); err != nil {
		panic(err.Error())
	}
	if err := DB.Ping(); err != nil {
		panic(err.Error())
	}
	log.Println("Successfully Connected to DB...")
	return
}
