# Goを学ぶ

- [はじめてのGo―シンプルな言語仕様，型システム，並行処理](http://gihyo.jp/dev/feature/01/go_4beginners)

## Docs

- [基本文法](./docs/chapter02.md)
- [型システム 〜型を用いた安全なプログラミング〜](./docs/chapter03.md)

## Goでhello world

`hello.go`

```go
// package名`main`を宣言
package main

// packageをインポートする
import(
  "fmt"
)

func main() {
  fmt.Println("hello world")
}
```

## プログラムの実行

`go run`コマンドでプログラムを実行できる。

```bash
$ go run hello.go
hello world
```

プログラムが実行できれば、`hello world`が出力される。

## コンパイル

`go build`コマンドでプログラムをコンパイルできる。

```bash
$ go build hello.go
$ ls
hello hello.go
$ ./hello
hello world
```

コンパイルしたファイルを実行できれば、`hello world`が出力される。

>hello.goをコンパイルすると，hello（Windowsの場合はhello.exe）という実行形式のバイナリファイルが生成されます。ここでは64ビットのOS Xを用いてコンパイルしたため，生成された実行ファイルは，同じ64ビットのOS XであればGoがインストールされていなくても実行できます。

## フォーマット

Goでは標準のコーディング規約が決まっており、`go fmt`コマンドで自動整形できる。

```bash
$ go fmt hello.go
```

## ドキュメント

`go doc`コマンドで標準パッケージやサードパーティパッケージのドキュメンを確認できる。

```bash
$ godoc fmt
```

以下のコマンドを実行すれば、サーバーが起動し、`http://localhost:3000/`でブラウザの公式サイトと同様のインターフェイスでドキュメントを確認できる。

```bash
$ godoc -http=":3000"
```

## Goのプロジェクト構成とパッケージ

`myproject`というプロジェクトの中で`gosample`というパッケージを作成、そのパッケージを`main`パッケージから呼び出すように構成してみる。

### ディレクトリ構成

```
myproject
├── bin # go install時の格納先
├── pkg # 依存パッケージのオブジェクトファイル
└── src # プログラムのソースコード
```

### 環境変数GOPATHの指定

`myproject`ディレクトリのパスを`GOPATH`という環境変数に指定する。

```bash
$ cd myproject
$ export GOPATH=`pwd` # myprojectをGOPATHに登録
```

Makefileなどの構成ファイルはなしで、依存関係を解決してビルドしてくれる。

### パッケージの作成

パッケージを作成していく、Goでは1つのパッケージは1つのディレクトリに格納する。

#### gosampleパッケージ

`myproject/src/gosample/gosample.go`

```go
package gosample

var Message string = "hello gworld
```

#### mainパッケージ

`myproject/src/main/main.go`

```go
package main
import (
    "fmt"
    "gosample"
)

func main() {
    fmt.Println(gosample.Message) // hello world
}
```

`gosample`パッケージをインポートし、`main()`で`gosample.Message`を利用している。

先ほど、`myproject`ディレクトリを`GOPATH`に指定したため、`$GOPATH/src/gosample`のようにパスが解決され、`gosample`パッケージのインポートが成功する。

### ビルドと実行

#### 実行

正しく`GOPATH`が設定されていれば、`gosample`パッケージの場所が正しく解決されるので、`go run`コマンドで`main.go`を正しく実行できる。

```bash
$ cd $GOPATH/src/main # `myproject/src/main/`に移動する
$ go run main.go
hello world
```

#### ビルド

このプログラムをビルドして、1つの実行形式のファイルを生成してみる。

`go build`コマンドで、コマンドを実行した階層に実行ファイルを作ることもできるが、`go install`コマンドを実行すれば、生成されたファイルが`$GOPATH/bin`に自動的に格納される。

```
$ cd $GOPATH/src/main
$ go install
```

```
myproject
├── bin
│ └── main
├── pkg
└── src
    └── gosample
    │  └── gosample.go
    └── main
        └── main.go
```

## パッケージの公開

`gosample`パッケージを以下のようにGitHubで公開してみる。

https://github.com/soarflat-sandbox/gosample

### 公開したパッケージを利用する

`gosample`パッケージをインポートする記述を以下のように書き換える。

```go
package main

import (
    "fmt"

    "github.com/soarflat-sandbox/gosample"
)

func main() {
    fmt.Println(gosample.Message)
}
```

### 公開したパッケージを取得する

```bash
go get github.com/soarflat-sandbox/gosample
```

取得したパッケージは以下のように`go get`コマンドで指定したパス（`github.com/soarflat-sandbox/gosample`）と同じディレクトリ構成でプロジェクト内に展開される。

```
└── github.com
    └── soarflat-sandbox
        └── gosample
            └── gosample.go
```

再度`main.go`をビルドして実行し、正しく動作すれば成功。

