package employeestorage

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// this struct contains db context also methods to create, update, read, delete restaurant .....
type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{db: db}
}