package services

import (
	"github.com/kenta-kenta/go-intermediate-myapi/models"
	"github.com/kenta-kenta/go-intermediate-myapi/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
	// dbに接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// 1. repositories層の関数SelectArticleDetailで記事の詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 2. repositories層の関数SelectCommentListでコメント一覧を取得
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 3. 2で得たコメント一覧を、1で得たArticle構造体に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func PostArticleService(article models.Article) (models.Article, error) {
	// dbに接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	newArticle, err := repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	// dbに接続
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleArray, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleArray, nil
}

func PostNiceService(article models.Article) (models.Article, error) {
	// dbに接続
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	article = models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}
	return article, nil
}
