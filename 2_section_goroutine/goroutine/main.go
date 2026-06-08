package main

import (
	"fmt"
	"time"
)

// 並行処理で実行したい関数
func printMessage(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		// 処理が見えやすいように少し待機
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Goルーチンとして起動
	go printMessage("Goルーチン呼び出し")

	fmt.Println("プログラム修了")
}
