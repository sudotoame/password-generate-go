// Package files работает с локальными файлам json
package files

import (
	"fmt"
	"os"

	"demo/password/output"

	"github.com/fatih/color"
)

type JSONDB struct {
	fileName string
}

func NewJSONDB(name string) *JSONDB {
	return &JSONDB{
		fileName: name,
	}
}

func (db *JSONDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, fmt.Errorf("ERROR: READ_FILE")
	}
	return data, nil
}

func (db *JSONDB) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		output.PrintErrorSwitch(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintErrorSwitch(err)
		return
	}
	color.Green("Файл %s успешно обновлен.\n", db.fileName)
}
