package main

import (
	"fmt"
)

func main() {
	// バッファサイズ2のチャネルを作成
	// 受信側が準いなくても、2つまでならデータを溜められる
	ch := make(chan int, 2)

	// 1つ目のデータを送信（ブロックされない）
	ch <- 1
	fmt.Println("1つ目を送信完了")

	// ２つ目のデータを送信（ブロックされない）
	ch <- 2
	fmt.Println("2つ目を送信完了")

	// バッファから取り出して表示
	fmt.Println("受信：", <-ch)
	fmt.Println("受信：", <-ch)
}
