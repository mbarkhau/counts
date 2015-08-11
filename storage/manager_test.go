package storage

import (
	"bytes"
	"counts/config"
	"counts/utils"
	"os"
	"path/filepath"
	"testing"
)

func initTest() {
	os.Setenv("COUNTS_DATA_DIR", "/tmp/count_data")
	os.Setenv("COUNTS_INFO_DIR", "/tmp/count_info")
	path, err := os.Getwd()
	utils.PanicOnError(err)
	path = filepath.Dir(path)
	configPath := filepath.Join(path, "config/default.toml")
	os.Setenv("COUNTS_CONFIG", configPath)
}

func TestNoCounters(t *testing.T) {
	initTest()
	//FIXME: size of cache should be read from config
	m1 := newManager()
	m2 := newManager()
	m1.Create("marvel")
	data1 := []byte("wolverine")
	m1.SaveData("marvel", data1, 0)
	data2 := m2.LoadData("marvel", 0, 0)
	if bytes.Compare(data1, data2) != 0 {
		t.Error("Expected data2 == "+string(data1)+" got", data2)
	}
}

func TestGetAllInfo(t *testing.T) {
	conf := config.GetConfig()
	testFilePath := filepath.Join(conf.GetInfoDir(), "test.json")

	f, err := os.Create(testFilePath)
	defer os.Remove(testFilePath)
	if err != nil {
		t.Fatal("Couldn't create test file")
	}
	f.WriteString(`{
		"id": "test",
		"type": "immutable",
		"capacity": 12345
	}`)
	m := newManager()
	infoDatas := m.GetAllInfo()
	if len(infoDatas) != 1 {
		t.Fatal("Expected exactly one infoData")
	}

}
