package util

import (
	"encoding/json"
	"io"
	"myapp/componentTest/db"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
)

func ToJSON(t *testing.T, s string) (expected interface{}) {
	t.Helper()
	_ = json.Unmarshal([]byte(s), &expected)
	return
}

func JSONResponseFile(t *testing.T, fileName string) (expected interface{}) {
	return JSONFile(t, fileName, "response")
}

func RequestBytes(t *testing.T, fileName string) (expected []byte) {
	t.Helper()
	jsonFile, _ := os.Open(filepath.Join(".", "request", fileName))
	bytes, _ := io.ReadAll(jsonFile)
	return bytes
}

func ExecSQL(t *testing.T, sql string) {
	t.Helper()
	db.GetDB().Exec(sql)
}

func JSONFile(t *testing.T, fileName string, folder string) (expected interface{}) {
	t.Helper()
	jsonFile, _ := os.Open(filepath.Join(".", folder, fileName))
	bytes, _ := io.ReadAll(jsonFile)
	_ = json.Unmarshal(bytes, &expected)
	return
}
