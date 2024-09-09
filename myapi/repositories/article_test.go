package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/kenta-kenta/go-intermediate-myapi/models"
	"github.com/kenta-kenta/go-intermediate-myapi/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// テスト結果として期待する値を定義
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtitle1",
			expected: models.Article{
				ID:       1,
				Title:    "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum:  2,
			},
		}, {
			testTitle: "subtitle2",
			expected: models.Article{
				ID:       2,
				Title:    "2nd",
				Contents: "second blog post",
				UserName: "saki",
				NiceNum:  4,
			},
		},
	}

	// テスト対象となる関数を実行
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(db, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// 期待する値と実行結果を比べる
			if got.ID != test.expected.ID {
				t.Errorf("get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}
