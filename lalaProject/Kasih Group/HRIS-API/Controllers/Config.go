package Config

import (
	"log"

	//"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

var err error

func Config() (*gorm.DB, error) {
	dsn := "sqlserver://kp2021:9szNu6uMt8@10.160.1.19?database=HR"
	//dsn := "sqlserver://sa:Lalaproject18.@localhost?database=HR"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	//db, err := gorm.Open("mssql", "sqlserver://kp2021:9szNu6uMt8@10.160.1.19?database=HR")
	if err != nil {
		log.Println("Connection Failed")
	} else {
		log.Println("Connection Success")
	}
	// db.AutoMigrate(&M.M_User{})
	return db, err
}
