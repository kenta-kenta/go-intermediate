package main

import (
	"database/sql"
	"dbsample/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//接続に使うユーザー・パスワード・データベース名を定義
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"

	//データベースに接続するためのアドレス文を定義
	//ここでは"docker:docker@tcp(127.0.0.1:3307)/sampledb?parseTime=true"となる
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	//Open関数を用いてデータベースに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}

	//プログラムが終了するときに、コネクションがcloseされるようにする
	defer db.Close()

	//クエリの定義
	articleID := 0
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`

	//クエリの実行
	row := db.QueryRow(sqlStr, articleID)
	//データ読み出しが0件だった場合
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	//rowsからのデータの読み出し
	//articleが複数入るスライスを作成する
	var article models.Article
	var createdTime sql.NullTime

	articleArray := make([]models.Article, 0)
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", articleArray)
}
