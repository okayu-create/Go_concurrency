package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// ループ変数を引数として渡す
		go func(n int) {
			fmt.Printf("Goルーチン番号：%d\n", n)
		}(i)
	}

	// すべてのGoルーチンが終わるのを待つ
	time.Sleep(500 * time.Millisecond)
}
