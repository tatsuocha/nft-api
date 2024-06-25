package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type CommonDao struct{}

func NewCommonDao() CommonDao {
	return CommonDao{}
}

func (dao CommonDao) Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=password port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースの接続が失敗しました。: ", err)
	}
	// データベース接続の確認
	log.Println("データベースの接続が完了しました。")

	return db
}
