package postgresql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type CommonDao struct{}

func NewCommonDao() CommonDao {
	return CommonDao{}
}

func (dao CommonDao) Connect() *gorm.DB {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_NAME") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースの接続が失敗しました。: ", err)
	}
	// データベース接続の確認
	log.Println("データベースの接続が完了しました。")

	return db
}
