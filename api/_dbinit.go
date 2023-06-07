package utils

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DbInit() {
	db, err := sqlx.Open("mysql", os.Getenv("DSN"))

}
