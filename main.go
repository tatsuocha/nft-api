package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nft-api/wire"
)

func main() {
	voiceHandler, err := wire.InitializeWire()
	if err != nil {
		log.Fatalf("依存性の注入に失敗しました。 %v", err)
	}

	router := gin.Default()
	voiceHandler.RegisterVoiceRoutes(router)
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("アプリケーションの起動に失敗しました。 %v", err)
	}
}
