package initializers

import (
	"fmt"
	"os"

	"github.com/vineet12344/jwtgolang/errorhandlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dns := os.Getenv("DB")
	var err error

	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	errorhandlers.CheckError(err)

	fmt.Printf("Database name is %s \n", DB.Name())

	fmt.Println("✅ Connected to PostgreSQL via GORM")

}
