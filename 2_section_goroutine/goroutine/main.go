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

	// 通常の関数呼び出し（この処理が終わるまで次は実行されない）
	printMessage("通常呼び出し")

	// Goルーチンの完了を待つための待機
	// ※本来、待機は4章で学ぶWaitGroupなどを使用する
	time.Sleep(1 * time.Second)
	fmt.Println("プログラム修了")
}
