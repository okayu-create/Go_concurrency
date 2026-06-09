package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("開始")

	// 無名関数をGoルーチンとして起動
	go func(msg string) {
		fmt.Println(msg)
	}("無名関数からのメッセージ")

	// 動作確認のための待機時間
	time.Sleep(100 * time.Millisecond)

	fmt.Println("終了")
}
