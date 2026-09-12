package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WaitGroupの変数を生成
	var wg sync.WaitGroup

	fmt.Println("メイン処理を開始")

	// 3つのGoルーチンを起動して並行処理
	for i := 0; i < 3; i++ {
		// Goルーチン起動前にカウンタを1つ増やす
		wg.Add(1)

		go func(id int) {
			// 関数終了時に必ずDone()を呼ぶようdefer文を使う
			defer wg.Done()

			fmt.Printf("ワーカー %d 作業開始\n", id)
			// 処理の進行を見やすくするために１秒待機
			time.Sleep(1 * time.Second)
			fmt.Printf("ワーカー %d 作業完了\n", id)
		}(i)
	}

	// 全てのGoルーチンがDoneメソッドを呼ぶ（カウンタが0になる）までここで待機
	fmt.Println("全ワーカーの完了を待機中...")
	wg.Wait()

	fmt.Println("すべての処理が完了しました。プログラムを終了します")
}
