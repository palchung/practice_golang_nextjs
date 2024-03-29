package gpostgres

import (
	"fmt"
	"root/gleam/golang/tool/setting"
)

func Dsn() string {

	host := setting.DatabaseSetting.Host
	port := setting.DatabaseSetting.Port
	user := setting.DatabaseSetting.User
	password := setting.DatabaseSetting.Password
	dbname := setting.DatabaseSetting.DBname
	s := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", host, port, dbname, user, password)

	return s
}
