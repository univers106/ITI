package config

import (
	_ "embed"
	"errors"
	"log/slog"
	"os"

	"go.yaml.in/yaml/v4"
)

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
