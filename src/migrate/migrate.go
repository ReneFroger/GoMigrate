package migrate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const MigrationsPath string = "./migrates"
const UpMigrationsPath string = "./migrates/up"
const DownMigrationsPath string = "./migrates/down"
const DatabaseConfigFilePath = "./config/database.json"

type DatabaseConfig struct {
	DriverName      string `json:"driver_name"`
	DataScourceName string `json:"data_source_name"`
}

func NewMigrate(name string) {
	prefix := time.Now().UTC().Format("19920709213000")

	downName := UpMigrationsPath + prefix + name + ".sql"
	os.Create(downName)
	upName := DownMigrationsPath + prefix + name + ".sql"
	os.Create(upName)
}

func Migrate() {
	databaseConfig := loadConfig(DatabaseConfigFilePath)

	db, err := sql.Open(databaseConfig.DriverName, databaseConfig.DataScourceName)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	filePathes, _ := filepath.Glob(UpMigrationsPath + "/*.sql")
	for _, filePath := range filePathes {
		execWithFile(db, filePath)
	}
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

func execWithFile(db *sql.DB, filePath string) {
	fmt.Println(filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	contentStr := string(content)

	tx, _ := db.Begin()
	_, err = tx.Exec(contentStr)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
}

func init() {
	os.Mkdir(MigrationsPath, os.ModePerm)
	os.Mkdir(UpMigrationsPath, os.ModePerm)
	os.Mkdir(DownMigrationsPath, os.ModePerm)
}
