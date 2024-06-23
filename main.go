package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nft-api/wire"
)

func main() {
	router := gin.Default()
	voiceHandler, _ := wire.InitializeVoiceHandler()
	ethereumHandler, _ := wire.InitializeEthereumHandler()
	voiceHandler.RegisterVoiceRoutes(router)
	ethereumHandler.RegisterEthereumRoutes(router)
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("アプリケーションの起動に失敗しました。 %v", err)
	}
}
