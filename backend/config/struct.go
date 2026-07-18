package config

type Config struct {
	Host        string `yaml:"host"`
	PostgresURL string `yaml:"postgres_url"`
}
