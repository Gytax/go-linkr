package config

type Config struct {
	Port   string
	DbPath string
}

func NewConfig(port string, dbPath string) *Config {
	return &Config{
		Port:   port,
		DbPath: dbPath,
	}
}
