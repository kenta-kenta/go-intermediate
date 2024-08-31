package main

import (
	"database/sql"
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

	//トランザクションの実行------------------------------------------------------------
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	//現在のいいね数を取得するクエリを実行する
	article_id := 1
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	//変数nicenumに現在のいいね数を読み込む
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	//いいね数を+1する更新処理を行う
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()

	// //データを挿入する処理-----------------------------------------------------------
	// article := models.Article{
	// 	Title:    "insert test",
	// 	Contents: "Can I insert data correctly?",
	// 	UserName: "saki",
	// }
	// const sqlStr = `
	// 	insert into articles (title, contents, username, nice, created_at)
	// 	values (?, ?, ?, 0, now())
	// `

	// result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())

	//データの読み出し---------------------------------------------------------
	// //クエリの定義
	// articleID := 0
	// const sqlStr = `
	// 	select *
	// 	from articles
	// 	where article_id = ?;
	// `

	// //クエリの実行
	// row := db.QueryRow(sqlStr, articleID)
	// //データ読み出しが0件だった場合
	// if err := row.Err(); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //rowsからのデータの読み出し
	// //articleが複数入るスライスを作成する
	// var article models.Article
	// var createdTime sql.NullTime

	// articleArray := make([]models.Article, 0)
	// err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// if createdTime.Valid {
	// 	article.CreatedAt = createdTime.Time
	// }

	// fmt.Printf("%+v\n", articleArray)
}
