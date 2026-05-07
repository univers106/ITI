package config

import (
	"crypto/rand"
	"os"

	"go.yaml.in/yaml/v4"
)

const (
	defaultDirPerm  os.FileMode = 0o750
	defaultFilePerm os.FileMode = 0o600
)

func getExampleConfig() []byte {
	var example Config

	var data []byte

	var err error

	example = Config{
		SessionKey: rand.Text(),
		Domain:     "localhost",
		DataDir:    "./data",
		Port:       8080,
		Address:    "localhost",
	}

	data, err = yaml.Marshal(example)
	if err != nil {
		panic("cannot marshal example config")
	}

	return data
}

func createExample(path string) []byte {
	data := getExampleConfig()

	err := os.WriteFile(path, data, defaultFilePerm)
	if err != nil {
		panic(err)
	}

	return data
}
