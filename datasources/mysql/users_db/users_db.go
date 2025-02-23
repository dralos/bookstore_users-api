package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

const (
	MYSQL_USERS_USERNAME = "MYSQL_USERS_USERNAME"
	MYSQL_USERS_PASSWORD = "MYSQL_USERS_PASSWORD"
	MYSQL_USERS_HOST     = "MYSQL_USERS_HOST"
	MYSQL_USERS_SCHEMA   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB
)

func SetupDatabase() error {
	username := os.Getenv(MYSQL_USERS_USERNAME)
	password := os.Getenv(MYSQL_USERS_PASSWORD)
	host := os.Getenv(MYSQL_USERS_HOST)
	schema := os.Getenv(MYSQL_USERS_SCHEMA)

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		schema,
	)

	client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		return err
	}
	if err := client.Ping(); err != nil {
		return err
	}

	Client = client

	mysql.SetLogger(log.New(log.Writer(), "", 0))

	log.Println("database successfully configured")
	return nil
}
