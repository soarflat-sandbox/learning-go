package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// チャネルを返す関数
// 戻り値を <-chan string と指定することで読み出し専用チャネルになる
// そのため、この関数の呼び出す側でチャネルを書き込むことを防ぐ
func getStatus(urls []string) <-chan string {
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
	return statusChan
}

// for/select 文と break を用いて実装したタイムアウト処理
func main() {
	// １秒後に値が読み出されるチャネル
	timeout := time.After(time.Second)
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}
	statusChan := getStatus(urls)

LOOP: // for/selet を抜けたい場合はラベル名（任意）が必要
	for {
		select {
		// statusChan からデータが読み出された場合
		case status := <-statusChan:
			fmt.Println(status) // 受信したデータを表示
		// timeout からデータが読み出された場合
		// 今回は１秒後にデータが読み出されているので、HTTP リクエストが完了しているかどうかに
		// 関わらず、1秒経過したらループを抜けてプログラムが終了する。
		case <-timeout:
			break LOOP // このfor/selectを抜ける
		}
	}
}
