package services

import "database/sql"

// サービス構造体の作成
type MyAppService struct {
	db *sql.DB
}

func NewAppService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
