package file_based

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	defaultDirPerm  os.FileMode = 0o750
	defaultFilePerm os.FileMode = 0o600
)

func createDirIfNotExists(dirPath string) {
	dir, err := os.Stat(dirPath)
	if !os.IsNotExist(err) {
		if !dir.IsDir() {
			panic(dirPath + " is exist and is not a directory")
		}

		return
	}

	err = os.Mkdir(dirPath, defaultDirPerm)
	if err != nil {
		panic(err)
	}
}

func saveStructToJsonFile(filePath string, data any) {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filePath, jsonData, defaultFilePerm)
	if err != nil {
		panic(err)
	}
}

func loadStructFromJsonFile[T any](filePath string) (T, error) {
	dir, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		if dir.IsDir() {
			return *new(T), os.ErrExist
		}
	}
	// #nosec G304
	data, err := os.ReadFile(filePath)
	if err != nil {
		return *new(T), fmt.Errorf("failed to read JSON file: %w", err)
	}

	var result T

	err = json.Unmarshal(data, &result)
	if err != nil {
		return *new(T), fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return result, nil
}
