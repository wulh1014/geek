package model

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
	Age  int
	Phone string
}

var db *gorm.DB

func GetSqlDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	return db
}

func GetUserByName(phone string) (user User, err error) {
	if err = GetSqlDB().Where("phone=?", phone).First(&user).Error; err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound){
			err = errors.Wrap(err,fmt.Sprintf("there is no %s record in database",phone))
			return
		}
		err = errors.Wrap(err,"internalErr")
		return
	}
	return
}
