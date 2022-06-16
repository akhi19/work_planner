package configs

import "os"

type SqlConfig struct {
	User     string
	Password string
	Port     string
	Database string
}

type Configuration struct {
	Port      string
	SqlConfig SqlConfig
}

var config Configuration

func init() {
	config = Configuration{
		Port: "8080",
		SqlConfig: SqlConfig{
			User:     os.Getenv("MSSQL_DB_USER"),
			Password: os.Getenv("MSSQL_DB_PASSWORD"),
			Port:     os.Getenv("MSSQL_DB_PORT"),
			Database: os.Getenv("MSSQL_DB_NAME"),
		},
	}
}
