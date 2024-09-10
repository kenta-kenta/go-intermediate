package services

import (
	"github.com/kenta-kenta/go-intermediate-myapi/models"
	"github.com/kenta-kenta/go-intermediate-myapi/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	// dbに接続する
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
