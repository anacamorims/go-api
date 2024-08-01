package router

import (
	"github.com/gin-gonic/gin"
)



func Initialize(){
	//iniciando router
	router := gin.Default()
	
	//iniciando routes
	initializeRoutes(router)


	//rodando a API
	router.Run(":8080")
}


