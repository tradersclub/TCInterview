package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tradersclub/TCInterview/entiti"
)

func Process_user(p string, u string) bool {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db-golang?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s")

	if err != nil {
		panic("failed to connect database")
	}

	var user []entiti.User

	result := db.Table("users").Where("email = ?", u).Find(&user)

	if result.Error != nil {
		return false
	} else {
		return true
	}
}

func CREATEUSER(u entiti.User) bool {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db-golang?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s")

	if err != nil {
		panic("failed to connect database")
	}

	result := db.Table("users").Create(&u)

	if result.Error != nil {
		return false
	} else {
		return true
	}
}

func GetUser(id int) (bool, entiti.User) {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db-golang?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s")

	if err != nil {
		panic("failed to connect database")
	}

	var user entiti.User

	result := db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return false, user
	} else {
		return true, user
	}
}
