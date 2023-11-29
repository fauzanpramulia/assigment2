package config

import (
	"fmt"
	"os"
	"strconv"
	"assigment2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var db *gorm.DB
var err error

var host = os.Getenv("PGHOST")
var port = os.Getenv("PGPORT")
var user = os.Getenv("PGUSER")
var pass = os.Getenv("PGPASSWORD")
var dbname = os.Getenv("PGDBNAME")

func ConnectGorm(){
	portDB, err:= strconv.Atoi(port)
	if err != nil{
		panic(err)
	}

	psqlInfo := fmt.Sprintf(`
	host=%s
	port=%d
	user=%s`+`
	password=%s
	dbname=%s
	sslmode=disable`, host, portDB, user, pass, dbname)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	db.AutoMigrate(models.Order{}, models.Item{})
	fmt.Println("Berhasil Terhubung Ke Database")
}

func GetDB() *gorm.DB {
	return db
}