package main

import (
	"fmt"
	"log"
	"net/http"
)

// チャネルを返す関数
// 戻り値を <-chan string と指定することで読み出し専用チャネルになる
// そのため、この関数の呼び出す側でチャネルを書き込むことを防ぐ
func getStatus(urls []string) <-chan string {
	// string を扱うチャネルの作成。チャネルを利用することで、ゴルーチン間でのデータのやりとりが可能
	// 今回は chan string を指定しているので string 型のデータの書き込みと読み出しができる。
	// また、第二引数に len(urls)（今回は配列が３つなので 3 になる）を渡しているので、
	// 同時に３つまではチャネル内部に値を保持できる。そのため３つまでの読み書きはブロックしない。
	// バッファをつけることで main() で実行しているチャネルの読み出しを待たずにゴルーチンを終了できる。
	statusChan := make(chan string, len(urls))
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

func main() {
	urls := []string{
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
