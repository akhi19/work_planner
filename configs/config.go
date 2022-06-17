package configs

type SqlConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

type Configuration struct {
	Port        string
	LogFilePath string
	SqlConfig   SqlConfig
}

var config Configuration

func init() {
	config = Configuration{
		Port: "8080",
		SqlConfig: SqlConfig{
			Host:     "postgres",
			User:     "postgres",
			Password: "postgres",
			Port:     "5432",
			Database: "planner",
		},
		LogFilePath: "./log",
	}
}

func GetConfig() Configuration {
	return config
}
