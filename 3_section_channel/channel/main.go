package main

import (
	"fmt"
)

func main() {
	// int型のデータを扱うチャネルを作成
	ch := make(chan int)

	// 新しいgoルーチンを起動してデータを送信
	go func() {
		fmt.Println("データを送信します")
		ch <- 100 // チャネルへ100を送信
	}()

	// チャネルからデータを受信
	// データがくるまでここで待機する
	result := <-ch

	fmt.Printf("受信した値：%d\n", result)
}
