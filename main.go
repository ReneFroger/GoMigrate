package main

import (
	_ "github.com/go-sql-driver/mysql"
	"migrate"
)

func main() {
	migrate.Migrate()
}
