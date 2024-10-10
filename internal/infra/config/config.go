package config

type Config struct {
	Port int
}

var (
	cfg *Config
)

func Get() *Config {
	return cfg
}

func Load() {
	cfg = &Config{
		Port: 8789,
	}
}
