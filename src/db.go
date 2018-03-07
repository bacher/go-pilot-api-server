package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Connect() {
	_db, err := sqlx.Connect("mysql", "root:@/team")

	if err != nil {
		panic(fmt.Errorf("Connection to MySQL failed: %s", err))
	}

	DB = _db
}
