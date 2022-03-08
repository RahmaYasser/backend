package driver

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/harranali/authority"
)

// GetDbConnection is used to connect to postgres database
func GetDbConnection() (*gorm.DB, *sql.DB) {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("HOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"))
	sqlDB, err := sql.Open("pgx", dataSourceName)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")
	log.Println("pinged db")
	//SeedDB(sqlDB)
	return db, sqlDB
}
func InitAuthority(db *gorm.DB) *authority.Authority {
	return authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           db,
	})
}
