package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

var DSN = "host=localhost user=root password=mysecretpassword dbname=banc-api port=3535"

func DBconection() {
	var err error

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("DB connection")
	}

}
