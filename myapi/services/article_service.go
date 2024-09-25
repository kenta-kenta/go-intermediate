package services

import (
	"database/sql"
	"errors"

	"github.com/kenta-kenta/go-intermediate-myapi/apperrors"
	"github.com/kenta-kenta/go-intermediate-myapi/models"
	"github.com/kenta-kenta/go-intermediate-myapi/repositories"
)

// サービス構造体のメソッド化
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	// article型とerror型を同時に扱う構造体
	type articleResult struct {
		article models.Article
		err     error
	}
	// articleResult型のチャネルを定義
	articleChan := make(chan articleResult)
	defer close(articleChan)

	// articleChanを通じて、SelectArticleDetail関数の結果を送信
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(s.db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(s.db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	// 1. repositories層の関数SelectArticleDetailで記事の詳細を取得
	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	// 2. repositories層の関数SelectCommentListでコメント一覧を取得
	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	// 3. 2で得たコメント一覧を、1で得たArticle構造体に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
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
