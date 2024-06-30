package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nft-api/app/domain/entity"
	"nft-api/wire"
	"os"
)

func main() {
	initDb()
	run()
}

func initDb() {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_NAME") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースの接続が失敗しました。: ", err)
	}
	// データベース接続の確認
	log.Println("データベースの接続が完了しました。")

	err = db.AutoMigrate(&entity.VoiceEntity{})
	if err != nil {
		log.Fatal("データベースの初期化に失敗しました。: ", err)
	}
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
