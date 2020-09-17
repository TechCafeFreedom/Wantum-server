package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"wantum/pkg/constants"
	"wantum/pkg/tlog"
)

func ConnectLocalDB() (*sql.DB, error) {

	dbuser := os.Getenv("MYSQL_USER")
	if dbuser == "" {
		dbuser = constants.MysqlDefaultUser
	}
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	if dbpassword == "" {
		dbpassword = constants.MysqlDefaultPassword
	}
	dbhost := os.Getenv("MYSQL_HOST")
	if dbhost == "" {
		dbhost = constants.MysqlDefaultHost
	}
	dbport := os.Getenv("MYSQL_PORT")
	if dbport == "" {
		dbport = constants.MysqlDefaultPort
	}
	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		dbname = constants.MysqlDefaultName
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbport, dbname)
	tlog.GetAppLogger().Info(fmt.Sprintf("connect db: %s", dataSource))

	return sql.Open("mysql", dataSource)
}