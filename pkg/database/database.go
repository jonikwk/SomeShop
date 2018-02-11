package database

import (
	"database/sql"
	"fmt"

	c "../configuration"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

//GetConnectionString -> Функция получения строки подключения базы данных
func GetConnectionString(config *c.Configuration) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Connect.DBHost, config.Connect.Port, config.User.Login, config.User.Password, config.Connect.DBName)
}

//OpenDB -> Функция открытия подклбчения базы данных
func OpenDB(config *c.Configuration, psqlInfo string) *sql.DB {
	var db = new(sql.DB)
	db, err := sql.Open(config.Connect.DBType, psqlInfo)
	if err != nil {
		glog.Exit(err)
	}
	return db
}
