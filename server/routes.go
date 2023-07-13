package server

import (
	"kaspin/server/handler"

	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server, reqLog, resLog, errLog zerolog.Logger) {
	transactionHandler := handler.NewTransactionHandler(reqLog, resLog, errLog)
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.engine.POST("/register", transactionHandler.Register())
	server.engine.POST("/status", transactionHandler.Status())
	server.engine.POST("/payment", transactionHandler.Payment())
	server.engine.POST("/callback", transactionHandler.Callback())
}
