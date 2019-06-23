# JSON/HTMLサーバ

POSTで送信されたJSONデータをファイルに保存し、リクエストに応じてファイルから読み出したデータをHTMLに格納して返すHTTPサーバ。

## Usage 

```shell
$ go run main.go
```

### POST

`id`の名前のファイルを作成し、`name`の値を書き込む。

```shell
curl http://localhost:3000/persons -d '{"id":1,"name":"gopher"}' 
```

### GET

ブラウザでアクセスすると、クエリに該当したデータがHTML上に出力される。

http://localhost:3000/persons?id=1