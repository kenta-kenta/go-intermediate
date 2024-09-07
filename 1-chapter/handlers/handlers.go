package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//1.リクエスト内容から、どのHTTPメソッドが使われているのかを判断する
//2.所定のメソッドが使われていた場合、正常応答を返す
//3.所定のメソッドが使われていなかった場合、エラーを返す

// reqは受け取る内容、wは書き込む内容みたいなイメージ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	//クエリパラメータを取得
	queryMap := req.URL.Query()

	var page int
	//pageが一つ以上あるとき
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		//pageの一つ目を採用する
		page, err = strconv.Atoi(p[0])

		//数値に変換できなかった時エラーを返す
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
