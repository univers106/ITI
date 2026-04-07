package config

import "go.yaml.in/yaml/v4"

func getExampleConfig(mock any) []byte {
	var example Config
	var data []byte
	var err error

	if mock != nil {
		data, err = yaml.Marshal(mock)
	} else {
		example = Config{
			JwtKey: "secret",
		}
		data, err = yaml.Marshal(example)
	}
	if err != nil {
		panic("cannot marshal example config")
	}
	return data

}
