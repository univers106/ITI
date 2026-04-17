package config

import (
	"os"

	"github.com/gorilla/securecookie"
	"go.yaml.in/yaml/v4"
)

func getExampleConfig() []byte {
	var example Config
	var data []byte
	var err error

	example = Config{
		SessionKey: string(securecookie.GenerateRandomKey(32)),
		Domain:     "localhost",
		DataDir:    "./data",
	}

	data, err = yaml.Marshal(example)
	if err != nil {
		panic("cannot marshal example config")
	}
	return data

}

func createExample(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := getExampleConfig()
	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
