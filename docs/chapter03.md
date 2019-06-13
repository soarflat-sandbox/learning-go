# 型システム 〜型を用いた安全なプログラミング〜

独自の型の宣言方法と使い方について。

## type

typeを利用すれば、既存の型を拡張した独自の型を定義できる。

```go
type ID int
type Priority int

func ProcessTask(id ID, priority Priority) {
}

var id ID = 3
var priority Priority = 5

// 第1引数にはID型の値を渡さなくてはいけないが
// 以下はPriority型の値を渡しているためコンパイルエラーになる
ProcessTask(priority, id)
```

### なぜtypeを利用するのか

厳密な型チェックが可能になり、堅牢なプログラムを記述できるから。

上記の例を`type`を利用せずに書くと以下のようなミスをする可能性がある。

```go
// 引数がint型であれば、なんでも受け入れる
func ProcessTask(id, priority int) {
}

id := 3 // int型に推論
priority := 5 // int型に推論
// 引数を渡す順番を間違えているが、どちらもint型なのでコンパイルエラーは発生しない。
ProcessTask(priority, id)
```

また、IDEのサポートも得られやすくなり、リファクタリング時のリグレッションも防ぎやすくなる。

## 構造体（struct）

フィールドの集まり。

以下は3つのフィールドを持つTask型を定義している。

```go
type Task struct {
  ID int
  Detail string
  done bool
}
```

フィールドの可視性は名前で決まる。大文字で始まるフィールドはパブリックになり、小文字で始まるフィールドはパッケージ内で閉じたスコープになる。

### フィールドに値を割りあてる

以下のようにすればフィールドに値を割り当てられる。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{
		ID:     1,
		Detail: "buy the milk",
		done:   true,
	}
	// 型推論使えるので以下のようにも書ける
	// task := Task{
	// 	ID:     1,
	// 	Detail: "buy the milk",
	// 	done:   true,
	// }
	fmt.Println(task.ID)     // 1
	fmt.Println(task.Detail) // "buy the milk"
	fmt.Println(task.done)   // true
}
```

### フィールド名の省略

構造体に定義した順で値を渡すことで、フィールド名を省略できる。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	var task Task = Task{1, "buy the milk", true}
	// 型推論使えるので以下のようにも書ける
	// task := Task{1, "buy the milk", true}
	fmt.Println(task.ID)     // 1
	fmt.Println(task.Detail) // "buy the milk"
	fmt.Println(task.done)   // true
}
```

### 構造体の生成時に値を明示的に指定しなかった場合

ゼロ値で初期化される。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func main() {
	task := Task{}
	fmt.Println(task.ID)     // 0
	fmt.Println(task.Detail) // ""
	fmt.Println(task.done)   // false
}
```

### ポインタ型

構造体をポインタ型で扱うことができる。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

var task Task = Task{} // Task型
var task *Task = &Task{} // Taskのポインタ型型（渡されたアドレスに存在する値の型はTask型）
```

#### ポインタ型の利用例

以下はポインタ型を利用していない例。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func Finish(task Task) {
  task.done = true
}

func main() {
  task := Task{done: false}
  Finish(task)
  fmt.Println(task.done) // false
}
```

呼び出し元の`task.done`は変更されていない。

場合によっては呼び出し元も変更したいこともある。

以下のようにポインタ型を定義し、ポインタを渡すことで呼び出し元も変更できる。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

// Taskポインタ型のポインタ変数を受け取る
func Finish(task *Task) {
  task.done = true
}

func main() {
  // ↓は`var task *Task = &Task{done: false}`と同じ
  task := &Task{done: false} // Taskポインタ型のポインタ変数taskを定義
  Finish(task)
  fmt.Println(task.done) // true
}
```

### new()

`new()`を利用すれば、構造体のフィールドを全てゼロ値で初期化し、そのポインタを返す。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func Finish(task *Task) {
	task.done = true
}

func main() {
	// ↓は`var task *Task = new(Task)
  task := new(Task)
  task.Detail = "buy the milk"
  Finish(task)

  fmt.Println(task.ID) // 0
  fmt.Println(task.Detail) // "buy the milk"
	fmt.Println(task.done) // true
}
```

### コンストラクタ

Goには構造体のコンストラクタにあたる構文がない。

代わりにNewで始まる関数を定義し、その内部で構造体を生成するのが通例。

たとえば、`Task`という構造体を生成する関数は`NewTask()`という関数にし、構造体を生成してポインタを返すようにする。

```go
func NewTask(id int, detail string) *Task {
  task := &Task{
    ID: id,
    Detail: detail,
    done: false,
  }
  return task
}

func main() {
  task := NewTask(1, "buy the milk")
  fmt.Printf("%+v", task) // &{ID:1 Detail:buy the milk done:false}
}
```

## メソッド

特定の型に関連付けられた（定義された）関数。

以下はメソッドの利用例。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

// taskという名前のTask型のレシーバーを持つことを意味する
// Task型にString()メソッドが定義される
func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func main() {
	task := NewTask(1, "buy the milk")
	fmt.Printf("%s", task.String()) // 1) buy the milk
}
```

上記の記述だと、Task型のコピーがレシーバーとして渡されるため、呼び出し元は変更されない。

```go
type Calc struct{ value, value2 int }

// 関数
func Add(q Calc) int {
	return q.value + q.value2
}

// Calc型のレシーバーを受け取るので、Calc型でこのメソッドが利用できるようになる
func (p Calc) Add() int {
	return p.value + p.value2
}

// Calc型のコピーがレシーバーとして渡されるため、呼び出し元は変更されない
func (p Calc) increment() int {
	p.value = p.value + 1
}

func main() {
	q := Calc{3, 2}     // 3 + 2 = 5
	fmt.Println(Add(q)) // 5

	p := Calc{3, 2}      // 3 + 2 = 5
  fmt.Println(p.Add()) // 5
  p.increment()
  fmt.Println(p.value) // 3
}
```

### レシーバーをポインタとして渡す

呼び出しもとに変更を反映させたい場合は、以下のようにレシーバーをポインタとして渡すようにする。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

// Taskのポインタがレシーバーとして渡されるため、呼び出し元が変更される
func (task *Task) Finish() {
	task.done = true
}

func main() {
	task := new(Task)
	task.Finish()
	fmt.Println(task.done) // true
}
```

## インターフェース

メソッドの定義された型。

構造体と同様で`type`の後ろに記述をする。

### インターフェースの宣言

以下はstring型を返す`String()`メソッドが定義されたインターフェース。

```go
type Stringer interface {
  String() string
}
```

インターフェースの名前は、定義したメソッド名が単純な場合、そのメソッド名に「er」を加えた名前をつける慣習がある。

そのため、`String()`が定義されたインターフェースの名前は`Stringer`になる。

### インターフェースの実装

Goでは、Javaのimplements構文のように、インターフェースを実装していることを明示的に宣言する構文はない。

型がインターフェースに定義されたメソッドを実装していれば、インターフェースを満たしているとみなす。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

type Stringer interface {
	String() string
}

// Stringer型（String()メソッドが実装された型）を引数にとる
func Print(stringer Stringer) {
	fmt.Println(stringer.String())
}

func main() {
  task := NewTask(1, "buy the milk")
  // taskにはString()メソッドが実装されているため、Print()に渡すことができる
	Print(task) // 1) buy the milk
}
```

### interface{}

以下はメソッドを定義していないインターフェース。

```go
type Any interface {
}
```

つまり、すべての型はこのインターフェースを実装していることになる。

これを利用すれば以下のようにどんな型も受け取ることができる関数を定義できる。

```go
func Do (e Any) {
  // do something
}

// 以下のようにも書ける
func DO (e interface{}) {

}
```

## 型の埋め込み

Goでは継承がサポートされていない。

代わりに他の型を「埋め込む」という方式で構造体やインターフェースの振る舞いを拡張する。

### 構造体の埋め込み

以下の`Task`構造体に対して、`User`構造体を埋め込んでみる。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
}
```

`User`構造体の定義は以下の通り。

`FulName()`メソッドとコンスラクタ関数も定義している。

```go
type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullName := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullName
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}
```

これを`Task`に埋め込むには、以下のように`Task`構造体の型宣言時にフィールドではなく、型のみを記述すれば良い。

```go
type Task struct {
	ID     int
	Detail string
	done   bool
	*User  // User構造体を埋め込む
}

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullName := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullName
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func NewTask(id int, detail, firstName, lastName string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}
	return task
}

func main() {
	task := NewTask(1, "buy the milk", "soar", "flat")
	// TaskにUserのフィールドが埋め込まれている
	fmt.Println(task.FirstName) // soar
	fmt.Println(task.LastName) // flat
	// TaskにUserのメソッドが埋め込まれている
	fmt.Println(task.FullName()) // soar flat
	// Taskから埋め込まれたUser自体にもアクセス可能
	fmt.Println(task.User) // ${soar flat}
}
```

### インターフェースの埋め込み

インターフェースも埋め込み可能であり、主な用途は複数のインターフェースを埋め込んで新たなインターフェースを定義することである。

たとえば、組み込みのioパッケージでは、ReaderやWriterなどのインターフェースが定義されている。

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}

type Writer interface {
  Write(p []byte) (n int, err error)
}
```

上記を以下のように埋め込むことで、両方のインターフェースが定義されたインターフェースを定義できる。

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}

type Writer interface {
  Write(p []byte) (n int, err error)
}

type ReadWriter interface {
  Reader
  Writer
}
```

## 型の変換

Goでは、暗黙的な型変換が起こることはないが、型を明示的に変換することはできる。

### キャスト（変換）

以下のように型をキャストできる。

```go
var i uint8 = 3
var j uint32 = uint32(i) // jにはiをuint8からuint32の型に変換した値が格納される
fmt.Println(j)           // 3

var s string = "abc"
var b []byte = []byte(s) // jにはsをstringから[]byteに変換した値が格納される
fmt.Println(b)           // [97 98 99]

// 以下はキャストに失敗する。キャストに失敗した場合はパニックが発生する
// a := int("a")
```

### Type Assertion（型の検査）

```go
func Print(value interface{}) {
  // valueがstring型であるか判定
  // 第１戻り値には、判定が成功した場合にその型に変換された値が返る
  // 第２戻り値には、判定が成功したかどうかの真偽値が返る
  s, ok := value.(string)
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string\n")
	}
}

func main() {
	Print("abc") // value is string: abc
	Print(10)    // value is not string
}
```

### Type Switch（型による分岐）

Type Assertionは単一の型に対する検査しかできないが、Type Switchを利用する複数の型に対する検査をできる。

```go
type Stringer interface {
	String() string
}

type Task struct {
	ID     int
	Detail string
	done   bool
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func Print(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %d\n", v)
	case Stringer:
		fmt.Printf("value is Stringer: %s\n", v)
	}
}

func main() {
	task := NewTask(1, "buy the milk")

	Print("abc") // value is string: abc
	Print(10)    // value is int: 10
	Print(task)  //value is Stringer: 1) buy the milk
}
```