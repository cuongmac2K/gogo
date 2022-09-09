package db

import (
	"first_demo/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config, errFile := config.ReadConfig("config/config.yml")
	if errFile != nil {
		panic("invalid ConfigPath")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBconf.User, config.DBconf.Pass, config.DBconf.Host, config.DBconf.Port, config.DBconf.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		//logs.Error(err)
		fmt.Println("CONNECT MYSQL FAILED!")
		fmt.Println(err.Error())
	}
	DB = db
	fmt.Println("CONNECT MYSQL SUCCESS!")

}
