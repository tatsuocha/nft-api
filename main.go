package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nft-api/wire"
)

func main() {
	connectPostgres()
	run()
}

func connectPostgres() {
	dsn := "host=localhost user=postgres password=password port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースの接続が失敗しました。: ", err)
	}
	// データベース接続の確認
	log.Println("データベースの接続が完了しました。")
}

func run() {
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
