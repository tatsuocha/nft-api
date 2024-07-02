package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nft-api/app/domain/entity"
	"nft-api/app/utils"
	"nft-api/wire"
)

func main() {
	migrateDb()
	runApplication()
}

func migrateDb() {
	dsn := "host=" + utils.GetEnv("POSTGRES_HOST") + " user=" + utils.GetEnv("POSTGRES_USER") + " password=" + utils.GetEnv("POSTGRES_PASSWORD") + " dbname=" + utils.GetEnv("POSTGRES_NAME") + " port=" + utils.GetEnv("POSTGRES_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
	}
	// データベース接続の確認
	log.Println("データベースの接続が完了しました。")

	err = db.AutoMigrate(&entity.VoiceEntity{})
	if err != nil {
		log.Fatalf("データベースのマイグレーションに失敗しました。: %v", err)
	}
}

func runApplication() {
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
