package config

import (
	_ "embed"
	"errors"
	"log/slog"
	"os"

	"go.yaml.in/yaml/v4"
)

func ReadConfig(path string) Config {
	// #nosec G304
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			slog.Warn("config.yaml not exist, creating example...")

			data = createExample(path)
		} else {
			panic("something wrong with config.yaml file: " + err.Error())
		}
	}

	answer := Config{}

	err = yaml.Unmarshal(data, &answer)
	if err != nil {
		panic("failed to unmarshal config: " + err.Error())
	}

	return answer
}
