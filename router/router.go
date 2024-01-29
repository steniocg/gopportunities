package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	//Run the server
	router.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
