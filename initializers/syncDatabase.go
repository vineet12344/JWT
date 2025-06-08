package initializers

import "github.com/vineet12344/jwtgolang/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
