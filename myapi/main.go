package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/kenta-kenta/go-intermediate-myapi/api"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	// APIを動かすのに必要なデータベースを用意する
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	// サーバを起動する
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
