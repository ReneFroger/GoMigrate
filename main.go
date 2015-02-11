package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"migrate"
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

	// migrate.NewMigrate("test2")
	// migrate.Rollback(db)
	// migrate.Migrate(db)
}

func init() {
	databaseConfig := loadConfig(DatabaseConfigFilePath)

	var err error
	db, err = sql.Open(databaseConfig.DriverName, databaseConfig.DataScourceName)
	if err != nil {
		panic(err)
	}
}
