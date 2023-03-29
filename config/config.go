package config

type Config struct {
	Server struct {
		Address string
	}
	Database struct {
		Server   string
		Port     int
		User     string
		Password string
		Database string
	}
}

func LoadConfig() *Config {

	cfg := &Config{}
	cfg.Server.Address = ":8080"
	cfg.Database.Server = "localhost"
	cfg.Database.Port = 1433
	cfg.Database.User = "sa"
	cfg.Database.Password = "MyPass@word"
	cfg.Database.Database = "GO"

	return cfg
}
