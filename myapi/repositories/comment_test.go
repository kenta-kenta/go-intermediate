package repositories_test

import (
	"testing"

	"github.com/kenta-kenta/go-intermediate-myapi/models"
	"github.com/kenta-kenta/go-intermediate-myapi/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2

	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Error(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test",
	}

	expectedCommentNum := 3

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("wants %d but got %d", expectedCommentNum, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?;
		`

		testDB.Exec(sqlStr, comment.Message)
	})
}
