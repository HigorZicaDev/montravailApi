package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações padrão
	router := gin.Default()
	// Defina uma rota para fazer requisição HTTP
	initializeRoutes(router)
	// Executa o Router passando a porta da Api
	err := router.Run(":8888")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
