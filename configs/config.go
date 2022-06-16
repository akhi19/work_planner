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
			Host:     "localhost",
			User:     "test",
			Password: "123",
			Port:     "55000",
			Database: "planner",
		},
		LogFilePath: "./log",
	}
}

func GetConfig() Configuration {
	return config
}
