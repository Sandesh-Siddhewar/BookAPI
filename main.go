package main

import (
	"github.com/Sandesh-Siddhewar/eBook/BOOKAPI/internal/database"
	"github.com/Sandesh-Siddhewar/eBook/BOOKAPI/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Database()

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8081")

}
