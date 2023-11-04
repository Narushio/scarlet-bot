package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func IsExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadJson[T any](filePath string) (T, error) {
	var data T
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func WriteJson[T any](filePath string, data T) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, b, 0644)
}
