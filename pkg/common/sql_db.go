package common

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/akhi19/work_planner/configs"
)

// Using mutex to avoid duplicate connections in case of parallel or concurrent calls
var lockMutex sync.Mutex
var sqlHandler *sql.DB

func InitializeConnection(config configs.Configuration) {
	lockMutex.Lock()
	defer lockMutex.Unlock()
	if sqlHandler == nil {
		user := config.SqlConfig.User
		password := config.SqlConfig.Password
		database := config.SqlConfig.Database
		port := config.SqlConfig.Port
		connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)
		handler, connectionError := sql.Open("mssql", connectionString)
		if connectionError != nil {
			//Panic : since no point starting application
			panic(fmt.Errorf("error opening database: %v", connectionError))
		}
		sqlHandler = handler
	}
}

func GetSqlHandler() *sql.DB {
	return sqlHandler
}
