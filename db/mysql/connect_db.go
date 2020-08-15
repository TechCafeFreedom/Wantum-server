package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"wantum/pkg/tlog"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	_ "github.com/go-sql-driver/mysql"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type mysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db"`
	Instance string `json:"instance"`
}

// CreateSQLInstance sqlのコネクションを作成
func CreateSQLInstance() *sql.DB {
	// create client(secret manager)
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		// connect local db
		return connectLocalSQL()
	}
	return connectCloudSQL(client, &ctx)
}

// connectLocalSQL localのmysqlのコネクション作成
func connectLocalSQL() *sql.DB {
	tlog.GetAppLogger().Debug("connectDB: local")
	dbuser := os.Getenv("MYSQL_USER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("MYSQL_HOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbport := os.Getenv("MYSQL_PORT")
	if dbport == "" {
		dbport = "3306"
	}
	dbname := os.Getenv("MYSQL_DATABASE")
	if dbname == "" {
		dbname = "wantum"
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbuser, dbpassword, dbhost, dbport, dbname)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		tlog.GetAppLogger().Panic(err.Error())
	}
	return db
}

// connectCloudSQL cloudSQLのコネクション作成
func connectCloudSQL(client *secretmanager.Client, ctx *context.Context) *sql.DB {
	tlog.GetAppLogger().Debug("connectDB: cloudSQL")
	projectID := "wantum-server"
	secretID := "mysql-config"
	// requestの作成
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, secretID),
	}

	// get secret value
	result, err := client.AccessSecretVersion(*ctx, accessRequest)
	if err != nil {
		tlog.GetAppLogger().Panic(fmt.Sprintf("failed to access secret version: %v", err))
	}

	// decode json
	var config mysqlConfig
	err = json.Unmarshal(result.Payload.Data, &config)
	if err != nil {
		tlog.GetAppLogger().Panic(fmt.Sprintf("failed to marshal json: %v", err))
	}

	// connect db
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s)/%s", config.User, config.Password, config.Protocol, config.Instance, config.DBName))
	if err != nil {
		tlog.GetAppLogger().Panic(fmt.Sprintf("failed to open sql: %v", err))
	}
	return db
}
