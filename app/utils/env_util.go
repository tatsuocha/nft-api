package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		log.Fatal("環境変数の取得に失敗しました。key：", key)
	}
	return env
}
