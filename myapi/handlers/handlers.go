package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kenta-kenta/go-intermediate-myapi/models"
	"github.com/kenta-kenta/go-intermediate-myapi/services"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

// GET /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	//パスパラメータ
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// service層の呼び出し
	article, err := services.GetArticleService(articleID)
	if err != nil {
		log.Printf("Error in ArticleDetailHandler: %v", err)
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	//jsonへのエンコード
	json.NewEncoder(w).Encode(article)
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	//json→goのデコード
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// service層の呼び出し
	article, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	//go→jsonのエンコード
	json.NewEncoder(w).Encode(article)
}

// GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	//クエリパラメータ
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])

		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	// sesrvice層の呼び出し
	articleList, err := services.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	//jsonへのエンコード
	json.NewEncoder(w).Encode(articleList)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// service層の呼び出し
	article, err := services.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	//jsonへのエンコード
	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// service層の呼び出し
	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	//jsonへのエンコード
	json.NewEncoder(w).Encode(comment)
}
