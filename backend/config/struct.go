package config

type Config struct {
	SessionKey string `yaml:"session_key"`
	Domain     string `yaml:"domain"`
	DataDir    string `yaml:"data_dir"`
}
