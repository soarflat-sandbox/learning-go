# 標準パッケージ 〜JSON、ファイル、HTTP、HTMLを扱う〜

## encoding/jsonパッケージ

### 構造体をJSONに変換

以下はencoding/jsonパッケージを利用して、構造体をJSONに変換している。

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
  }
  // ポインタを渡せばJSON文字列の[]byteに変換される
  // ポインタじゃなくても変換された
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
  }
  // byteを文字列に変換する
	fmt.Println(string(b)) // {"ID":1,"Name":"Gopher","Email":"gopher@example.org","Age":5,"Address":""}
}
```

変換されたJSONは、キーの名前は構造体のフィールド名と同じになっている。

また、小文字で始まるプライベートなフィールドはJSONに含まれていない。

プライベートではないフィールドも出力しないようにしたり、出力されるキーの名前を変えたい場合などは、構造体にタグを記述することで出力をコントロールできる。

encoding/jsonパッケージで利用できるタグは以下のようなものがある。

```go
`json:"name"`       // nameというキーで格納する
`json:"-"`          // JSONに格納しない
`json:",omitempty"` // 値が空なら無視
`json:",string"`    // 値をJSONとして格納
```

タグは以下のように型定義の後ろに記述する。

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID      int `json:"id"`
	Name    string
	Email   string
	Age     int `json:"-"`
	Address string
	memo    string `json:",string"`
}

func main() {
	person := &Person{
		ID:      1,
		Name:    "Gopher",
		Email:   "gopher@example.org",
		Age:     5,
		Address: "",
		memo:    "golang lover",
	}

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b)) // {"id":1,"Name":"Gopher","Email":"gopher@example.org","Address":""}
}
```

### JSONから構造体へ変換

```go
type Person struct {
	ID      int
	Name    string
	Email   string
	Age     int
	Address string
	memo    string
}

func main() {
	var person Person
	b := []byte(`{"id":1,"name":"Gopher","age":5}`)
	err := json.Unmarshal(b, &person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person) // {1 Gopher 5 }
}
```

## os、ioパッケージ

### ファイルの生成

```go
package main

import (
	"log"
	"os"
)

func main() {
	// ファイルを生成
	// ファイル名を渡すと*os.File構造体へのポインタを取得できる
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// プログラムの実行が終了したらファイルを閉じる
	defer file.Close()
}
```

`*os.File`は`io.ReadWriteCloser`という複数のインターフェースが定義されたインターフェース型。

このインターフェースには`Read()`、`Writre()`、`Close()`の3つのメソッドが実装されている。

### ファイルへの書き込み

`*os.File`には`io.Writer`インターフェースが実装されており、以下のように定義されている。

```go
type Writer interface {
  Write(p []byte) (n int, err error)
}
```

[]byte型の値を引数として渡すと、その中身を対象に書き込み、戻り値として書き込んだバイト数とエラーを返す。

以下は指定したファイルにhello worldを書き込む例。

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// 書き込むデータを[]byteで用意する
  message := []byte("hello world\n")
  // WriteStringを利用すれば[]byteに変換する必要はない
  // message :=  file.WriteString("hello world\n")
  

	// 書き込みをする
	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
```

## net/httpパッケージ

以下はnet/httpパッケージを利用して、hello worldを返す簡単なサーバを実装している例。

```go
package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
```