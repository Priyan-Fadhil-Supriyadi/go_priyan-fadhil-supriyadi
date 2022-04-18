package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT string
	DB_HOST string
	DB_NAME string
}

func InitConfiguration() Config {
	return Config{
		DB_USERNAME: "root",
<<<<<<< HEAD
		DB_PASSWORD: "root",
		DB_NAME: "bimo",
		DB_PORT: "3306",
		DB_HOST: "0.0.0.0",
=======
		DB_PASSWORD: "",
		DB_NAME: "docker",
		DB_PORT: "3306",
		DB_HOST: "127.0.0.1",
>>>>>>> 2bbdc36abb92c50f1f5e872d14d94c34757f4f37
	}
}