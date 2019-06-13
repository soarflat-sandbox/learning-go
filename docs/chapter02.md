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

エラーを戻り値で表現できない場合は、`panic`と`recovery`を利用する。



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
func fn(arr [4]string) {
	arr[0] = "x"
	fmt.Println(arr) // [x b c d]
}

func main() {
	arr := [4]string{"a", "b", "c", "d"}
	fn(arr)
	fmt.Println(arr) // [a b c d]
}
```

## 配列

Goの配列は固定長。

以下は長さ4で要素の型stringである配列。

```go
var arr1 [4]string
```

アクセスは他の言語同様にインデックスでアクセスする。

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

配列の型は長さも情報として含むため、以下の`arr1`と`arr2`は、要素の型は同じstringだが長さが違うため配列としては別の型になる。

```go
var arr1 [4]string
var arr2 [5]string
```

そのため、以下のような`[4]string`型を引数にとる関数へ型のあわない`arr2`を渡すとコンパイルエラーになる。

```go
func fn(arr [4]string) {
	fmt.Println(arr)
}

func main() {
	var arr1 [4]string
	var arr2 [5]string

	fn(arr1) // ok
	fn(arr2) // コンパイルエラー
}
```

### 関数に配列を渡す場合は値渡しになる

関数に配列を渡すと、配列のコピーが渡される。

そのため、もとの配列には変更は反映されない。

```go
func fn(arr [4]string) {
  arr[0] = "x"
  fmt.Println(arr) // [x b c d]
}

func main() {
	arr := [4]string{"a", "b", "c", "d"}
  fn(arr)
  // もとの配列には変更は反映されない
	fmt.Println(arr) // [a b c d]
}
```

### ポインタを渡すことで参照渡しもできる

```go
func fnP(arr *[4]int) {
  for i, _ := range arr {
    arr[i] = 0
  }
  fmt.Println(arr) // [0, 0, 0, 0]
}

func main() {
  arr := [4]int{1, 2, 3, 4}
  fnP(&arr)
  fmt.Println(arr) // [0, 0, 0, 0]
}
```

## スライス

可変長の配列のこと（JavaScriptで利用している配列はこっち）。

以下のように宣言する。

```go
var s []string
```

初期化を行う場合は、配列と同じように書ける。

```go
s := []string{"a", "b", "c", "d"}
fmt.Println(s[0]) // "a"
```

### append()

スライスの末尾に値を追加し、その結果を返す関数。

```go
var s []string
// 追加した結果を返す
s = append(s, "a")
s = append(s, "b")
// 複数の値も追加できる
s = append(s, "c", "d")
fmt.Println(s) // [a b c d]
```

以下のようにスライスに別のスライスの値を展開できる。

```go
s1 := []string {"a", "b"}
s2 := []string {"c", "d"}
s1 = append(s1, s2)
fmt.Println(s1) // [a b c d]
```

### range

配列やスライスの値を反復処理できる構文。`for`と併用する。

```go
var arr [4]string

arr[0] = "a"
arr[1] = "b"
arr[2] = "c"
arr[3] = "d"

for i, s := range arr {
	// i = インデックス, s = 値
	fmt.Println(i, s)
}
```

stringやマップなどに対しても利用できる。

### 値の切り出し

```go
s := []int{0, 1, 2, 3, 4, 5}

// インデックス2から4までを切り出す、終点は含まない
fmt.Println(s[2:4]) // [2 3]

// インデックス0からスライスの長さ文を切り出す（要は全て切り出す）
fmt.Println(s[0:len(s)]) // [0 1 2 3 4 5]

// 始点を省略しているため、先頭から切り出す
fmt.Println(s[:3]) // [0 1 2]

// 終点を省略しているため、末尾まで切り出す
fmt.Println(s[3:]) // [3 4 5]

// 全部
fmt.Println(s[:]) // [0 1 2 3 4 5]
```

### 可変長引数

引数を以下のように指定すると、可変長引数として任意の数の引数をその型のスライスとして受け取ることができる。

```go
// numsは[]int型になる
func sum(nums, ...int) (result int) {
  for _, n := range nums {
    result += n
  }
  return
}

func main() {
  fmt.Println(sum(1, 2, 3, 4)) // 10
}
```

上記の場合、`1, 2, 3, 4`を渡しており、int型のスライス（`[]int{1, 2, 3, 4}`）として受け取る。

## マップ

値をKey-Value（キーバリュー）型で保存するデータ構造（JavaScriptのオブジェクトみたいなやつ）。

### 宣言と初期化

以下はint型のキーにstring型の値を格納するマップの宣言。

```go
var month map[int]string = map[int]string{}
// 以下はNG？
// var month = map[int]string{}
```

以下のようにキーを指定して値を保存する。

```go
month[1] = "January"
month[2] = "February"
fmt.Println(month) // map[1:January 2:February]
```

宣言と初期化は以下のように記述する。

```go
month := map[int]string{
  1: "January",
  2: "February"  
}
```

### マップの操作

#### 値を取得する

キーを指定すれば、マップの値を取得できる。

```go
month := map[int]string{
  1: "January",
  2: "February",
}

jan := month[1]
fmt.Println(jan) // January
```

#### キーの存在をチェックする

2つめの戻り値も受け取るようにすると、指定したキーがこのマップに存在するかどうかをboolで返す。

```go
month := map[int]string{
  1: "January",
  2: "February",
}

_, ok := month[1]
if ok {
  // データがあった場合
}
```

#### 指定したキーのデータを削除する

マップからデータを消す場合は`delete()`を利用する。

```go
month := map[int]string{
  1: "January",
  2: "February",
}
delete(month, 1)
fmt.Println(month) //  map[1:January]
```

#### 反復処理をする

スライスと同様で、`range`を利用すればfor文で反復処理が可能。

しかし、マップの場合、処理の順番は保証されないため注意。

```go
month := map[int]string{
  1: "January",
  2: "February"  
}

for key, value := range month {
  fmt.Printf("%d %s\n", key, value)
}
// 1 January
// 2 February
```

## ポインタ

型の前に`+`をつけることで、ポインタ型を利用できる。

また、アドレスは以下のように変数の前に`&`をつけて取得できる。

```go
var i int = 10
fmt.Println(i)  // 10
fmt.Println(&i) // 0xc000016090
```

以下のようにポインタ型とアドレスを利用すれば、参照渡しができる。

```go
func callByValue(i int) {
  i = 20 // 値を上書きする
}

func callByRef(i *int) {
  i = 20 // 参照先を上書きする
}

func main() {
  var i int = 10
  callByValue(i) // 値を渡す
  fmt.Println(i) // 10

  callByValue(&i) // アドレスを渡す
  fmt.Println(i) // 20
}
```

## defer

関数が終了する前に必ず実行する処理を定義できる。

以下の場合、何がおこっても必ずファイルを閉じたいため、`defer`文で`file.Close()`を指定している。

```go
func main() {
  file, err := os.Open("./error.go")
  if err != nil {
    // エラー処理
  }
  // 関数を抜ける前に必ず実行される
  defer file.Close()
  // 正常処理
}
```

関数を途中で抜ける処理があったり、パニックが発生しても必ず実行したい処理を定義したい場合に利用する。

## パニック

配列やスライスの範囲外にアクセスした場合などでは、エラーを返すことができない（戻り値で表現できない）ため、代わりにパニックという方法でエラーが発生する。

### `recover()`

パニックで発生したエラーは`recover()`という組み込みの関数で取得できるため、以下のように`recover()`を`defer`文の中で定義すれば、パニックで発生したエラーを処理してから関数を抜けられる。

```go
func main() {
  defer func() {
    err := recover()
    if err != nil {
      // runtime error: index out of range
      log.Fatal(err)
    }
  }()

  a := []int{1, 2, 3}
  // スライスの範囲外にアクセスするため、パニックが発生する
  fmt.Println(a[10])
}
```

### `panic()`

`panic()`という組み込み関数を利用すれば、自分でパニックを発生させることができる。

```go
a := []int{1, 2, 3}
for i := 0; i < 10; i++ {
  if i >= len(a) {
    panic(errors.New("index out of range"))
  }
  fmt.Println(a[i])
}
// プログラムを実行すると以下の出力がされる
// 1
// 2
// 3
// panic: index out of range
```

#### パニックを利用するうえでの注意点

パニックを利用するのは

- エラーを戻り値として表現できない
- 回復不可能なシステムエラーが発生した
- やむを得ず大域脱出が必要
   
などの場合であり、基本的にエラーは関数の戻り値として返すようにする。

