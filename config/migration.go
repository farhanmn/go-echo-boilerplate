package config

import (
	"fmt"
	"go-echo-experiment/internal/model"
)

func RunMigration() {
	db := ConnectGORM()

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Migration failed:", err)
	} else {
		fmt.Println("Migration completed successfully!")
	}
}
