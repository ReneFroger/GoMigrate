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
	json_string := `
{
"driver_name" : "mysql",
"data_source_name" : "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
}
`
	databaseConfig := &DatabaseConfig{}

	json.Unmarshal([]byte(json_string), &databaseConfig)

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
