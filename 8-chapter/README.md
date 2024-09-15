## ユニットテスト(応用編)

### httpハンドラのテストを書こう
テスト作成の流れ
1. テストしたい関数で使うリソースを作成
2. テスト対象の関数に入れるinputを定義
3. テスト対象の関数を実行してoutputを得る
4. outputが期待通りかチェック

- req *http.Request: 受け取ったhttpリクエスト
    - NewRequest関数: ハンドラのhttp.Request型をGoのコード内で作ることができます。
- w http.ResponseWriter: httpレスポンスの内容をこれに書き込む
    - NewRecorder関数: httptest.ResponseRecoder構造体を作成する

### サービス層のモックを作ろう
- モックの役割
    今まではデータベースを起動する必要があった  
    モックとは「それを呼び出す側が本来欲しがっていた形のデータを、仮で返すためのもの」