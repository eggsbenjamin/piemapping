package repository

import (
	"database/sql"
	"fmt"

	"github.com/eggsbenjamin/piemapping/commons"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

//	construct db querystring
func getDBQueryString(params map[string]string) (qs string) {
	for k, v := range params {
		qs += fmt.Sprintf("%s=%s&", k, v)
	}
	return
}

//	construct db connection string
func getConnectionString(params map[string]string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		viper.GetString("db_user"),
		viper.GetString("db_password"),
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_name"),
		getDBQueryString(params),
	)
}

//	connects to the db and returns the connection object.
//	this function panics rather than errors due it's criticality.
func NewConnection(log commons.LevelledLogWriter, params map[string]string) (DB *sql.DB) {
	var (
		err     error
		DBType  = viper.GetString("db_type")
		connStr = getConnectionString(params)
	)
	log.Info("Connecting to DB...")
	if DB, err = sql.Open(DBType, connStr); err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}
	log.Info("Successfully Connected to DB.")
	return
}
