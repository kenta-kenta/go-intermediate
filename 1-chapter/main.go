package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenta-kenta/go-intermediate-chapter1/handlers"
)

func main() {
	//ルータを明示
	r := mux.NewRouter()
	//ハンドラを定義→パスで対応付ける

	//定義したハンドラを、サーバで使用するように登録
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")   //日時とともにログに出力される
	log.Fatal(http.ListenAndServe(":8080", r)) //ListenAndServeでサーバを起動 Fatalはエラーが出たらプログラムを終了させる

	// 2行で書いた場合
	// err := http.ListenAndServe(":8080", nil)
	// log.Fatal(err)
}

//まとめ
//1.リクエストメソッドやパスから、ルータが適切なハンドラに処理を任せる
//2.ハンドラが適切なレスポンスを作成する。その際にパスパラメータやクエリパラメータを利用できる。
