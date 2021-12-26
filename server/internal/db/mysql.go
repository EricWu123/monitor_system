package db

import (
	"database/sql"
	"monitor_system/config"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlHandler *sql.DB
	mysqlMutex   sync.Mutex
)

func NewMysql(conf *config.DBConfig) *sql.DB {
	if mysqlHandler != nil {
		return mysqlHandler
	}

	mysqlMutex.Lock()
	defer mysqlMutex.Unlock()

	// double check
	if mysqlHandler != nil {
		return mysqlHandler
	}
	dsn := conf.UserName + ":" + conf.Password +
		"@tcp(" + conf.Host + ")/" + conf.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlHandler, err := sql.Open("mysql", dsn)
	if err != nil || mysqlHandler == nil {
		return nil
	}

	return mysqlHandler
}
