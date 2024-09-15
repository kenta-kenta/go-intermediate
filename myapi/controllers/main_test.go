package controllers_test

import (
	"testing"

	"github.com/kenta-kenta/go-intermediate-myapi/controllers"
	"github.com/kenta-kenta/go-intermediate-myapi/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()             // 2. サービス構造体にはsql.DB型が必要
	aCon = controllers.NewArticleController(ser) // 1. コントローラにはサービス構造体が必要

	m.Run()
}
