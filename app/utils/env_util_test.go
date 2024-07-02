package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// 準備
	key := "TEST_KEY"
	value := "test_value"
	os.Setenv(key, value)

	// 実行&比較
	assert.Equal(t, value, GetEnv(key))
}
