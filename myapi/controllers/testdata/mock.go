package testdata

import "github.com/kenta-kenta/go-intermediate-myapi/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return aritcleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return aritcleTestData, nil
}

func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return aritcleTestData[0], nil
}

func (s *serviceMock) PostNiceService(article models.Article) (models.Article, error) {
	return aritcleTestData[0], nil
}

func (s *serviceMock) PostCommentService(article models.Article) (models.Article, error) {
	return aritcleTestData[0], nil
}
