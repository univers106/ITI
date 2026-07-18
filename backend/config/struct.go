package config

type Config struct {
	Host           string `yaml:"host"`
	PostgresSqlURL string `yaml:"postgre_sql_url"`
}
