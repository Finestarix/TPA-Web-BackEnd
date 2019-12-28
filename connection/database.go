package connection

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"github.com/lib/pq"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var databaseName, databasePassword, databaseUser, databaseHost, databasePort, databaseSSL, databaseType string
var connectionString string

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Failed to Load .env File")
		os.Exit(3)
	}

	databaseName = os.Getenv("databaseName")
	databasePassword = os.Getenv("databasePassword")
	databaseUser = os.Getenv("databaseUser")
	databaseHost = os.Getenv("databaseHost")
	databasePort = os.Getenv("databasePort")
	databaseSSL = os.Getenv("databaseSSL")

	databaseType = os.Getenv("databaseType")

	connectionString =  "dbname=" + databaseName + " " +
					    "password=" + databasePassword + " " +
						"user=" + databaseUser + " " +
						"host=" + databaseHost + " " +
						"port=" + databasePort + " " +
						"sslmode=" + databaseSSL
}

func GetConnection() *gorm.DB {
	database, error := gorm.Open(databaseType, connectionString)
	if error != nil {
		panic("Connection Failed !" + string(error.Error()))
	}
	return database
}