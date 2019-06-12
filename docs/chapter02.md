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