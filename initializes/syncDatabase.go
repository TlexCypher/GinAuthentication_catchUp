package initializes

import "github.com/TlexCypher/ginAuthenticationCatchUp/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
