package main

import "fmt"

func main() {
	// string型のデータをやり取りするチャネルを作成
	ch := make(chan string)

	// 新しいGoルーチンを起動
	go func() {
		// チャネルに文字列を送信
		ch <- "Hello, World"
	}()

	// チャネルからデータを受信して変数に格納
	msg := <-ch

	// 受信したデータを表示
	fmt.Println(msg)
}
