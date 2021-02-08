package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jsonhut?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "jsonhut:y8Wx4ZkMXjnAnMfz@tcp(127.0.0.1:3306)/jsonhut?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//err2 := db.AutoMigrate(&Json{})
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	return nil
}