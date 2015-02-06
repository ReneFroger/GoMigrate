package migrate

import (
	"testing"
)

func TestPreVersion(t *testing.T) {
	filePathes := []string{"1", "3", "2"}
	curVersion := "2"
	if preVersion(filePathes, curVersion) != "1" {
		t.Errorf("出错了")
	}

	filePathes = []string{}
	curVersion = "2"
	if preVersion(filePathes, curVersion) != "2" {
		t.Errorf("出错了")
	}

	filePathes = []string{"3"}
	curVersion = "2"
	if preVersion(filePathes, curVersion) != "2" {
		t.Errorf("出错了")
	}
}
