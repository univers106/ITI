package config

import (
	_ "embed"
	"errors"
	"log/slog"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	JwtKey string `yaml:"jwt_key"`
}

func ReadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Warn("config.yaml not exist, creating example...")
			data, err = createExample(path)
			if err != nil {
				panic("failed to create example config: " + err.Error())
			}

		} else {
			panic("something wrong with config.yaml file: " + err.Error())
		}
	}
	answer := Config{}
	if err := yaml.Unmarshal(data, &answer); err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}
	return answer
}

func createExample(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data := getExampleConfig(nil)
	written, err := file.Write(data)
	if err != nil {
		return nil, err
	}
	if written != len(data) {
		return nil, errors.New("partial write of example config")
	}
	return data, nil
}
