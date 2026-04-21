package main

import (
	"errors"
	"fmt"
)

// TODO(human): ここに divide 関数を書く
//   シグネチャ: func divide(a, b int) (int, error)
//   やること: a を b で割った結果を返す。ただし b が 0 の時はエラーを返す
//   ヒント:
//     - b == 0 なら errors.New("cannot divide by zero") でエラーを作って返す
//       そのとき1つ目の戻り値（int）は 0 でOK → return 0, errors.New(...)
//     - b != 0 なら計算して返す。エラー側は nil
//       → return a / b, nil
//     - 「エラーも1つの戻り値」という感覚を掴むのがゴール

func main() {
	// TODO(human): 3パターン divide を呼んで結果を表示する
	//   1. divide(10, 2) → 正常系（5 が返る）
	//   2. divide(10, 0) → エラー系（0除算）
	//   3. divide(7, 3)  → 正常系（2 が返る。int の切り捨て）
	//
	// 各呼び出しの後で if err != nil { ... } else { ... } でチェック
	//
	// 書き方のヒント（1ケースだけ例示）:
	//   result, err := divide(10, 2)
	//   if err != nil {
	//       fmt.Println("10 / 2 = エラー:", err)
	//   } else {
	//       fmt.Println("10 / 2 =", result)
	//   }
	//
	// 完成イメージ:
	//   10 / 2 = 5
	//   10 / 0 = エラー: cannot divide by zero
	//   7 / 3 = 2
}
