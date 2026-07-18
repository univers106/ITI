package config

type Config struct {
	Host           string `yaml:"host"`
	PostgresSqlURL string `yaml:"postgreSQL_url"`
}
