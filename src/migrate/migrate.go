package migrate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const (
	MigrationsPath          = "./migrates"
	UpMigrationsPath        = "./migrates/up"
	DownMigrationsPath      = "./migrates/down"
	DatabaseConfigFilePath  = "./config/database.json"
	DatabaseVersionFilePath = "./migrates/version"
)

type DatabaseConfig struct {
	DriverName      string `json:"driver_name"`
	DataScourceName string `json:"data_source_name"`
}

func NewMigrate(name string) {
	prefix := time.Now().UTC().Format("19920709213000")

	downName := UpMigrationsPath + prefix + "_" + name + ".sql"
	os.Create(downName)
	upName := DownMigrationsPath + prefix + "_" + name + ".sql"
	os.Create(upName)
}

func Migrate() {
	databaseConfig := loadConfig(DatabaseConfigFilePath)

	db, err := sql.Open(databaseConfig.DriverName, databaseConfig.DataScourceName)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	curVersion, err := ioutil.ReadFile(DatabaseVersionFilePath)
	if err != nil {
		os.Create(DatabaseVersionFilePath)
	}
	if len(curVersion) == 0 {
		curVersion = []byte("0")
	}

	curVersionNum, err := strconv.Atoi(string(curVersion))
	if err != nil {
		panic("Version file must store a number")
	}

	filePathes, _ := filepath.Glob(UpMigrationsPath + "/*.sql")
	regex, _ := regexp.Compile(`^(\d+)`)
	for _, filePath := range filePathes {
		fileVersion := regex.FindString(path.Base(filePath))
		fileVersionNum, _ := strconv.Atoi(fileVersion)
		if curVersionNum < fileVersionNum {
			execWithFile(db, filePath)
			curVersionNum = fileVersionNum
			ioutil.WriteFile(DatabaseVersionFilePath, []byte(fileVersion), os.ModePerm)
		}
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
