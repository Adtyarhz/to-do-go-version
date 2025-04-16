package main

import (
	"go-lang-p1/config"
	"go-lang-p1/models"
	"go-lang-p1/routes"
)

func main() {
    config.ConnectDB()
    config.DB.AutoMigrate(&models.Todo{})

    r := routes.SetupRouter()
    r.Run(":8080")
}
