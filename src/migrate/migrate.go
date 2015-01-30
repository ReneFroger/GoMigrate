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

type DatabaseConfig struct {
	DriverName      string `json:"driver_name"`
	DataScourceName string `json:"data_source_name"`
}

func NewMigrate(name string) {
	os.Mkdir("./migrates", os.ModePerm)
	os.Mkdir("./migrates/up", os.ModePerm)
	os.Mkdir("./migrates/down", os.ModePerm)

	prefix := time.Now().UTC().Format("19920709213000")

	downName := "./migrates/up/" + prefix + name + ".sql"
	os.Create(downName)
	upName := "./migrates/down/" + prefix + name + ".sql"
	os.Create(upName)
}

func Migrate() {
	databaseConfig := loadConfig("./config/database.json")

	db, err := sql.Open(databaseConfig.DriverName, databaseConfig.DataScourceName)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	filePathes, _ := filepath.Glob("./migrates/up/*.sql")
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
