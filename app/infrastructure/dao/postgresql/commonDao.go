package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"nft-api/app/utils"
)

type CommonDao struct{}

func NewCommonDao() CommonDao {
	return CommonDao{}
}

func (dao CommonDao) Connect() *gorm.DB {
	dsn := "host=" + utils.GetEnv("POSTGRES_HOST") + " user=" + utils.GetEnv("POSTGRES_USER") + " password=" + utils.GetEnv("POSTGRES_PASSWORD") + " dbname=" + utils.GetEnv("POSTGRES_NAME") + " port=" + utils.GetEnv("POSTGRES_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースの接続が失敗しました。: ", err)
	}
	// データベース接続の確認
	log.Println("データベースの接続が完了しました。")

	return db
}
