package filebased

import (
	"encoding/json"
	"os"
)

func createDirIfNotExists(dirPath string) {
	if dir, err := os.Stat(dirPath); !os.IsNotExist(err) {
		if !dir.IsDir() {
			panic(dirPath + " is exist and is not a directory")
		}
		return
	}
	if err := os.Mkdir(dirPath, 0755); err != nil {
		panic(err)
	}
}

func saveStructToJsonFile(filePath string, data any) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		panic(err)
	}
}

func loadStructFromJsonFile[T any](filePath string) (T, error) {
	if dir, err := os.Stat(filePath); !os.IsNotExist(err) {
		if dir.IsDir() {
			return *new(T), os.ErrExist
		}
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return *new(T), err
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return *new(T), err
	}
	return result, nil
}
