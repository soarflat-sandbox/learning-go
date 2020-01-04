package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	// 起動したすべてのゴルーチンの終了を待ち合わせるために利用する
	wait := new(sync.WaitGroup)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	for _, url := range urls {
		// waitGroup に追加
		wait.Add(1)
		// go というキーワードをつけると、それぞれの関数が別のゴルーチンで実行される
		// そのため、リクエスト処理が並列で実行される
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			// waitGroupから削除（カウントが減る）
			wait.Done()
		}(url)
	}
	// カウントが 0　になるまで待つ
	wait.Wait()
}
