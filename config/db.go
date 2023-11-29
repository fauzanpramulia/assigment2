package config

import (
	"fmt"
	"os"
	"assigment2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)


var db *gorm.DB
var err error
func ConnectGorm(){
	err := godotenv.Load()
	if err != nil{
		panic(err)
	}

	var(
		host 		= os.Getenv("PGHOST")
		port 		= os.Getenv("PGPORT")
		user 		= os.Getenv("PGUSER")
		password	= os.Getenv("PGPASSWORD")
		dbname		= os.Getenv("PGDBNAME")	
	)

	psqlInfo := fmt.Sprintf(`
	host=%s
	port=%s
	user=%s`+`
	password=%s
	dbname=%s
	sslmode=disable`,
	 host, port, user, password, dbname)

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