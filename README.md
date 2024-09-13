# APIを作りながら進むGo中級者への道
## 1.HTTPサーバー
1. ウェブサーバーのHello, World！
2. ハンドラを増やそう
3. 許可するHTTPメソッドを指定しよう
4. gorilla/muxパッケージを使ってみよう
5. パスパラメータを取得できるようになろう
6. クエリパラメータを取得できるようになろう

## 2.構造体とjsonの扱い方
1. 構造体を定義しよう
2. Go構造体をjsonにエンコードしよう
3. jsonフィールドをカスタマイズしよう
4. HTTPレスポンスボディにjsonを書き込む
5. HTTPリクエストからボディの中身を読み取る
6. jsonをGo構造体に変換する
7. デコーダ・エンコーダを使ってリファクタしよう

## 3.データベースの扱い方
1. 前準備～DBセットアップ
2. Goのコードからデータベースに接続してみよう
3. データ取得処理を書いてみよう
4. NULLかもしれない値を受け取る
5. クエリに変数を混ぜよう
6. データ挿入処理を実行しよう
7. トランザクションを利用した更新処理を実行してみよう
8. APIで使うデータベース操作を実装しよう

## 4.ユニットテスト(基礎編)
1. ユニットテストを書いてみよう
2. テーブルドリブンテストを実装しよう
3. テストの前処理・後処理を書こう
4. 個別のテストケースごとの後処理を書こう
5. repositriesパッケージのテストを完成させよう

## 5.サービス層の作成
1. サービス層の実装
2. サービス層を利用してAPIを完成させよう

## 6.アーキテクチャ大改装
1. サービス層を大改装
2. サービス層を使う側を大改装
3. ルータ層の作成
4. インターフェースによる抽象化・疎結合化
5. インターフェースの小型化による役割の分離

## 7.エラー処理
1. 独自エラー型を導入しよう
2. 独自エラー型への置き換え - サービス層編
3. 独自エラー型への置き換え - コントローラ層編
4. エラーを適切にハンドリングしよう
