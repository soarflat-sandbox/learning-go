package main

import (
	"fmt"
	"log"
	"net/http"
)

var empty struct{} // サイズがゼロの構造体

func getStatus(urls []string) <-chan string {
	// string を扱い、値を３つ保持できるチャネルの作成
	statusChan := make(chan string, 3)
	// 値（struct{} の構造体）を５つ保持できるチャネルを作成
	// そのため、今回はゴルーチンを同時に５個まで実行できる
	limit := make(chan struct{}, 5)
	go func() {
		for _, url := range urls {
			select {
			// limit にデータを書き込んだ場合（書き込めるといういうことはゴルーチンの実行数が５個に達していない）
			case limit <- empty:
				go func(url string) {
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					// statusChan に res.Status を書き込む
					statusChan <- res.Status
					// limit を呼び出すので、チャネルに保持できる値の空きができる
					<-limit
				}(url)
			}
		}
	}()
	return statusChan
}
func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
		"http://example.com",
		"http://example.net",
		"http://example.org",
		"http://example.com",
		"http://example.net",
		"http://example.org",
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	statusChan := getStatus(urls)

	// 今回の場合 fmt.Println(<-statusChan) を３回実行しているため、
	// statusChan に３回書き込みがされて、3回読み出すまで処理がブロックされる。
	// そのため、3回読み出すまで main() の処理は終了しない。
	for i := 0; i < len(urls); i++ {
		// チャネルから res.Status を読み出す
		// 読み出す順番は書き込んだ順から読み出す
		fmt.Println(<-statusChan)
	}
}
