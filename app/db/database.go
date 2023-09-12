package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "github.com/BelCattaneo/go-chat/app/model"
)

// Env holds database connection to Postgres
type Env struct {
	DB *sql.DB
}

// database variables
// usually we should get them from env like os.Getenv("variableName")
const (
	host     = "postgresdb"
	port     = 5432
	user     = "spuser"
	password = "SPuser96"
	dbname   = "postgres"
)

// ConnectDB tries to connect DB and on succcesful it returns
// DB connection string and nil error, otherwise return empty DB and the corresponding error.
func ConnectDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &gorm.DB{}, err
	}
	fmt.Println("connected to database")
	return db, nil
}

func CreateTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Room{})
}
