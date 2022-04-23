package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func InitConfiguration() Config {
	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_NAME:     "unit_test",
		DB_PORT:     "3306",
		DB_HOST:     "127.0.0.1",
	}
}
