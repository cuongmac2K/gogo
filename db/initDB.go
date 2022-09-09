package db

import "first_demo/model/user"

func init() {
	//create table user
	var u user.User
	DB.Model(&u).AutoMigrate()

}
