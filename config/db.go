package config

import (
	"echo/model"
	"fmt"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "employee"
)

var (
	// db  *sql.DB
	db  *gorm.DB
	err error
)

// func loadConfig() error {
// 	err = godotenv.Load()
// 	if err != nil {
// 		panic(err) //gabisa load env
// 	}

// 	err = env.
// }

// func init() {
// 	err := loadConfig()
// 	if err != nil {
// 		panic(err)
// 	}

// 	Connect()
// }

// func Connect() (*sql.DB, error) {
func Connect() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// db, err = sql.Open("postgres", psqlInfo)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&model.Employee{}, &model.Item{})

	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected to database")

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
