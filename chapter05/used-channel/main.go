package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	// string を扱うチャネルの作成。チャネルを利用することで、ゴルーチン間でのデータのやりとりが可能
	// 今回は chan string を指定しているので string 型のデータの書き込みと読み出しができる。
	statusChan := make(chan string)
	for _, url := range urls {
		// go というキーワードをつけると、それぞれの関数が別のゴルーチンで実行される
		// そのため、リクエスト処理が並列で実行される
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			// チャネルに res.Status を書き込む
			statusChan <- res.Status
		}(url)
	}
	// 今回の場合 fmt.Println(<-statusChan) を３回実行しているため、
	// statusChan に３回書き込みがされて、3回読み出すまで処理がブロックされる。
	// そのため、3回読み出すまで main() の処理は終了しない。
	for i := 0; i < len(urls); i++ {
		// チャネルから res.Status を読み出す
		// 読み出す順番は書き込んだ順から読み出す
		fmt.Println(<-statusChan)
	}
}
