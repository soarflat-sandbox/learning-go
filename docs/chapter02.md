# 基本文法

## mainパッケージ

プログラムをコンパイルして実行すると、まず`main`パッケージの中にある`main()`関数が実行されるため、ここに処理を記述する。

```go
package main

func main() {
}
```

## import

```go
import (
  "fmt"
)

func main() {
  fmt.Println("hello world")
}
```

複数のパッケージのimportも可能。

```go
import (
  "fmt"
  "github.com/wdpress/gosample"
  "strings"
)
```

Goの標準パッケージはインストールしたGoの中に含まれているため、自動でパスが解決される。

それ以外のパッケージを利用したい場合は、`GOPATH`環境変数を指定してパスを解決する必要がある。

### オプションの指定

importにはいくつかのオプションが指定できる。

```go
import (
  // fmtをfとして利用する
  f "fmt"
  // importしたパッケージを利用しないことをコンパイラに伝える
  _ "github.com/wdpress/gosample"
  // パッケージ名を省略する
  . "strings"
)

func main() {
  // fmt.Println()がf.Println()になり
  // strings.ToUpper()がToUpper()になっている
  // github.com/wdpress/gosampleを利用していないが、エラーは発生しない
  f.Println(ToUpper("hello world"))
}
```

## 変数

```go
var message string = "hello world"

func main() {
    fmt.Println(message)
}
```

- `var`: 変数の宣言
- `message`: 変数名
- `string`: 変数の型
- `"hello world"`: 変数の中身

### 1度に複数の宣言と初期化

```go
var foo, bar, buz string = "foo", "bar", "buz"

func main() {
  fmt.Println(foo) // -> "foo"
  fmt.Println(bar) // -> "bar"
  fmt.Println(buz) // -> "buz"
}
```

### 変数宣言のショートハンド

変数宣言と初期化を関数の内部で行う場合は、`:=`を利用することで、varと型宣言を省略できる。

```go
func main() {
  // var message string = "hello world"
  // 以下は↑と同じ意味
  message := "hello world"
  fmt.Println(message)
}
```

変数の型はコンパイラによって推論される。

今回、変数に文字列を代入しているため、`message`の型はstring型になる。

## 定数

```go
func main() {
  const Hello string = "hello"
  Hello = "bye" // cannot assign to Hello
}
```

error型は定数宣言できない。

## ゼロ値

変数を宣言し、明示的に値を初期化しなかった場合、変数はゼロ値というデフォルト値で初期化される。

ゼロ値は型ごとに決まっている。たとえばintのゼロ値は0であるため、以下のコードは0を出力する。

```go
func main() {
  var i int // iはゼロ値で初期化
  fmt.Println(i) // 0
}
```

## if

if文の条件部に丸括弧は必要ない。

```go
func main() {
  a, b := 10, 100
  if a > b {
    fmt.Println("a is larger than b")
  } else if a < b {
    fmt.Println("a is smaller than b")
  } else {
    fmt.Println("a equals b")
  }
}
```

波括弧は必須（ないとエラーが出る）。

また、参考演算子は存在しない。

## for

条件部に丸括弧は必要ない。

```go
func main() {
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
}
```

Goでは繰り返しを表現する方法はforしかなく、while文やdo/while文などは存在しない。

### break、continue

```go
func main() {
  n := 0
  for {
    n++
    if n > 10 {
      break // ループを抜ける
    }
    if n%2 == 0 {
      continue // 偶数なら次の繰り返しに移る
    }
    fmt.Println(n) // 奇数のみ表示
  }
}
```

## switch文

```go
func main() {
	n := 10
	switch n {
	case 15:
		fmt.Println("FizzBuzz")
	case 5, 10:
		fmt.Println("Buzz")
	case 3, 6, 9:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
}
```

JavaScriptなどでは各`case`毎に`break`を記述する必要があるが、Goでは不要。

### fallthrough

複数の`case`の処理を実行さてない場合は`fallthrough`を記述する。

```go
func main() {
	n := 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		fmt.Println(n) // 0
	}
}
```

### 式での分岐

`case`に値だけではない式も指定できるため、if/else文の代わりに利用できる。

```go
func main() {
	n := 10
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%5 == 0:
		fmt.Println("Buzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
}
```

## 関数

```go
func hello() {
	fmt.Println("hello")
}

func main() {
	hello() // hello
}
```

### 引数がある場合

引数の型を指定する必要がある。

```go
// int型の引数i,jをとる
func sum(i, j int) { // func sum(i int, j int) と同じ
	fmt.Println(i + j)
}

func main() {
	sum(1, 2) // 3
}
```

### 戻り値がある場合

引数の次に型を指定する。

```go
// int型の値を返す
func sum(i, j int) int {
	return i + j
}

func main() {
	n := sum(1, 2)
	fmt.Println(n) // 3
}
```

### 複数の値を返す

戻り値が複数の場合、以下の`(int, int)`のように、型をカンマで区切って指定し、丸括弧でくくる。

```go
func swap(i, j int) (int, int) { // int型の戻り値を2つ返す
  return j, i
}

func main() {
	x, y := 3, 4
  x, y = swap(x, y)
  fmt.Println(x, y) // 4, 3

  // 戻り値を格納する変数を必要な数だけ用意していなため、コンパイルエラーになる。
  x = swap(x, y) 

  // `_`を指定すれば、戻り値を無視できる。
  x, _ = swap(x, y)
  fmt.Println(x) // 3

  // コンパイル、実行ともに可能
  // これを利用したいシーンがわかっていない
  swap(x, y)
}
```

### エラーを返す

Goでは複数の値を返すことができるため、内部で発生したエラーも戻り値で表現する。

たとえば、ファイルをを開く`os.Open()`は、1つめの戻り値に`*os.File`を返し、2つめの戻り値に`error`を返す。

関数の処理に成功した場合、`error`は`nil`になり、異常があった場合は`error`に値が入る。異常があった場合、`error`以外の値はゼロ値になる。

```go
func main() {
	file, err := os.Open("hello.go")
	if err != nil {
		// エラー処理
		// returnなどで処理を抜ける
	}
	// 正常時の処理
}
```

#### 自作のエラーを返す

`errors`パッケージを利用して、自作のエラーを作成できる。

```go
package main

import (
	"errors"
	"fmt"
	"log"
)

func div(i, j int) (int, error) {
	if j == 0 {
		// 自作のエラーを返す
		return 0, errors.New("divided by zero")
	}
	return i / j, nil
}

func main() {
	n, err := div(10, 0)
	if err != nil {
		// エラーを出力しプログラムを終了する。
		log.Fatal(err)
	}
	fmt.Println(n)
}
```

複数の値を返す場合もエラーを最後に返す慣習があるため、自分でAPIを設計する場合もエラーを返すのは最後にした方が良い。

異常を戻り値で表現できない場合は、`panic`と`recovery`を利用する（後述）。

### 名前付き戻り値

Goでは、戻り値にあらかじめ名前をつけることができる。

```go
// `result`と`err`という名前の戻り値
func div(i, j int) (result int, err error)
```

戻り値に名前をつけている場合は、`return`のあとに返す値を明示する必要がなく、`return`された時点で名前付き戻り値の値が自動的に返される。

```go
func div(i, j int) (result int, err error) {
	if j == 0 {
		err = error.New("divided by zero")
	}
  result = i / j
  return // 名前付き戻り値が自動的に返されるため、`return result, nil`と同じ
}
```

名前付き戻り値を用いることで

- 関数の宣言から戻り値の意味が読み取りやすくなる
- 戻り値のための変数の初期化が不要になる
- 同じ型の戻り値が多かった場合の`return`の書き間違えなどを防げる

などのメリットがある。

戻り値に名前をつけた場合でも、`return`のあとに戻す値を明示することは可能なので必要に応じて使い分ける。

### 関数リテラル

関数リテラルをを利用すると、無名関数を作成できる。

以下のように記述することで、関数を即時に実行できる。

```go
func main() {
  func(i, j int) {
    fmt.Println(i + j)
  }(2, 4)
}
```

関数を変数に代入したり、関数を関数の引数に渡すこともできる。

```go
var sam = func(i, j int) {
  fmt.Println(i + j)
}

func main() {
  sum(2, 4)
}
```

## 配列

Goの配列は固定長。

以下は長さ4で要素の型stringである配列。

```go
var arr1 [4]string
```

アクセスは他の言語同様に添字でアクセスする。

```go
var arr1 [4]string

arr[0] = "a"
arr[1] = "b"
arr[2] = "c"
arr[3] = "d"
fmt.Println(arr[0]) // a
```

宣言と同時に初期化することも可能であり、`[...]`を利用すれば、必要な配列の長さを暗黙的に指定できる。

```go
// どちらも同じ意味
arr := [4]string{"a", "b", "c", "d"}
arr := [...]string{"a", "b", "c", "d"}
```

### 配列の型は長さも情報として含む