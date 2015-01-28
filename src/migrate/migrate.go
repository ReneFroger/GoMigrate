package migrate

import (
	"os"
	"time"
)

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

}
