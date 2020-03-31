package database_client

import (
	"fmt"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	UserName = "root"
	Password = "root"
	DBName   = "users"
	Port     = 5432
	HostName = "localhost"
)

func GetDatabaseInstance() *gorm.DB {

	dataSource := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", HostName, Port, UserName, DBName, Password)
	db, err := gorm.Open("postgres", dataSource)

	if err != nil {
		_ = db.Close()
		panic("database not working")
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