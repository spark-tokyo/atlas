package utils

import "time"

const (
	// 負荷のかかる処理は少ないためサーバーのタイムアウトは5分に設定
	RequestTimeout time.Duration = 5 * time.Minute
)
