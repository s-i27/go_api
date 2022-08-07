package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	// リトライ処理:Exponential Backoff
	// ref : https://qiita.com/taizo/items/c397dbfed7215969b0a5
	get()
}

func get() {
	url := "http://google.co.jp"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer access-token")

	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)

	client := new(http.Client)
	resp, _ := client.Do(req)

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	// POSTメソッド
// 	// url.Values{}でPOSTで送信する入れ物を準備
// 	ps := url.Values{}

// 	// Add()でPOSTで送信するデータを作成
// 	ps.Add("id", "1")
// 	ps.Add("name", "もりぴ")
// 	// 特殊文字や日本語をエンコード
// 	fmt.Println(ps.Encode())

// 	// http.PostForm()でPOSTメソッドを発行
// 	res, err := http.PostForm("http://example.com", ps)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// deferでクローズ処理
// 	defer res.Body.Close()
// 	// Bodyの内容を読み込む
// 	body, _ := io.ReadAll(res.Body)
// 	// Bodyの内容を出力する
// 	fmt.Print(string(body))

// 	fmt.Printf("Hello World\n")
// }
