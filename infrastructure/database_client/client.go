package database_client

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	envUserName = "POSTGRES_USER"
	envPassword = "POSTGRES_PASSWORD"
	envDBName   = "POSTGRES_DB"
	envPort     = "POSTGRES_PORT"
	envHostName = "POSTGRES_HOST"
)



func GetDatabaseInstance() *gorm.DB {
	
		userName := os.Getenv(envUserName)                       // "root"
		password := os.Getenv(envPassword)                       // "root"
		dBName   := os.Getenv(envDBName)                         // "users"
		port, _  := strconv.ParseInt(os.Getenv(envPort), 10, 64) // 5432
		hostName := os.Getenv(envHostName)                       // "postgres"
	
	dataSource := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", hostName, port, userName, dBName, password)

	db, err := gorm.Open("postgres", dataSource)
	if err != nil {
		_ = db.Close()
		panic(err)
	}

	db.SingularTable(true)
	migrateDatabase(db)

	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.UserDb{})
}

func getUsers(db *gorm.DB) {
	user := models.UserDb{}
	db.First(&user, 1233899201)
	fmt.Println(user.Name)
}
