# 並行処理

## Goにおける並行処理
並行：構成のこと  
並列：処理のこと
### go 文におけるゴール―チンの起動 - 基礎編
関数の前にgoをつけると別のゴール―チンの中で関数が実行されるようになる

## 並行処理ができそうな場所を探す
- ベンチマークの用意  
    時間を計測するための機能  
    テストコードを書く  
    -benchをつけてテスト実行

## ゴール―チンを使って並行処理を書いてみる - ロックと待ち合わせ編
### go文によるゴール―チンの起動 - 応用編
1. 無名関数を定義
2. 定義した無名関数を実行
3. それをgo文の後に置く

### 並行処理の「難しさ」
- Race Condition(競合状態)とは  
    それぞれの処理が行われるタイミングに依存して処理が変わりうる、特に意図せぬ結果が生まれてしまうような状態のこと
- Race Conditionを避けるためには
    sync.Mutexによるロックとsync.WaitGroupによる待ち合わせ

## ゴール―チンを使って並行処理を書いてみる - チャネル編
### チャネルによる値の送受信
- チャネルに値を送信するゴールーチンは、受信側が値を受信する準備ができるまで待ちの状態になる
- チャネルから値を送信するゴールーチンは、送信側が値を送信する準備ができるまで待ちの状態になる
