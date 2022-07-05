package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME     = "talent"
)

// kita buat koneksi dgn db posgres
func SetupModels() *gorm.DB {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open("postgres", dbinfo)

	// check the connection
	checkErr(err)

	db.AutoMigrate(&Talent{})
	return db
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
