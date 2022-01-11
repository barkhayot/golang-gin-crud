package database

import (
	"editt/models"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// For Database Conneciton

/*DB is connected database object*/
var DB *gorm.DB

func Setup() {
	host := "localhost"
	port := "5432"
	dbname := "editt"
	user := "postgres"
	password := "010203"

	db, err := gorm.Open("postgres",
			"host="+host+
				" port="+port+
				" user="+user+
				" dbname="+dbname+
				" sslmode=disable password="+password)

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate([]models.Post{})
	DB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}