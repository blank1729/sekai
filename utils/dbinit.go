package utils

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DbInit() {
	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		fmt.Println("unable to connect to db", err)
	}
	db.Ping()

}
