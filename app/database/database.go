package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/BelCattaneo/go-chat/app/model"
)

// database variables
// TODO: get them from env like os.Getenv("variableName")
const (
	host     = "postgresdb"
	port     = 5432
	user     = "spuser"
	password = "SPuser96"
	dbname   = "postgres"
)

// ConnectDB tries to connect DB and on succcesful it returns
// nil error, otherwise return corresponding error.
// it stores db conn in package exported variable
func ConnectDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &gorm.DB{}, err
	}
	fmt.Println("connected to database")
	return conn, nil
}

func CreateTables(conn *gorm.DB) {
	conn.AutoMigrate(&model.User{}, &model.Room{})
}

func SetupDB() {
	conn, _ := ConnectDB()
	CreateTables(conn)
	db, _ := conn.DB()
	defer db.Close()
}
