package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`   // idというキーで格納する
	Name string `json:"name"` // nameというキーで格納する
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

// テンプレートのコンパイル
var t = template.Must(template.ParseFiles("index.html"))

func personHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// POSTリクエストをJSONに変換し、
	// 変換した値からファイルを作成してその中にNameの値を書き込む
	if r.Method == "POST" {
		// リクエストボディをJSONに変換
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		// ファイル名を{id}.txtにする
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// ファイルにNameを書き込む
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// レスポンスとしてステータスコード201を送信
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		// パラメータを取得
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// personを生成
		person := Person{
			ID:   id,
			Name: string(b),
		}

		// レスポンスにエンコーディングしたHTMLを書き込む
		t.Execute(w, person)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/persons", personHandler)
	http.ListenAndServe(":3000", nil)
}
