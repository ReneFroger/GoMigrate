package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"migrate"
	"os"
)

const DatabaseConfigFilePath = "./config/database.json"

var db *sql.DB

type DatabaseConfig struct {
	DriverName      string `json:"driver_name"`
	DataScourceName string `json:"data_source_name"`
}

func loadConfig(filePath string) *DatabaseConfig {
	jsonChars, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	databaseConfig := &DatabaseConfig{}

	json.Unmarshal(jsonChars, &databaseConfig)

	return databaseConfig
}

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "install":
		os.OpenFile(DatabaseConfigFilePath, os.O_CREATE, 0644)
		migrate.Install()
	case "new":
		migrate.NewMigrate(flag.Arg(1))
	case "rollback":
		migrate.Rollback(db)
	case "refresh":
		migrate.RefreshSchema(db)
	default:
		migrate.Migrate(db)
	}
}

func init() {
	databaseConfig := loadConfig(DatabaseConfigFilePath)

	var err error
	db, err = sql.Open(databaseConfig.DriverName, databaseConfig.DataScourceName)
	if err != nil {
		panic(err)
	}
}
