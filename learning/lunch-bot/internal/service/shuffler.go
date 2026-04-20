package service

import (
	"math/rand"
)

// Shuffle は参加者IDをランダムに3〜5人組のグループに分割する。
//
// 入力例: ["U001", "U002", "U003", "U004", "U005", "U006", "U007"]
// 出力例: [["U003", "U005", "U001", "U007"], ["U002", "U006", "U004"]]
//
// ルール:
//   - 各グループは 3 人以上 5 人以下
//   - 入力人数が 3 未満のときは nil を返す（呼び出し側でフォールバック）
//   - 順序はランダム化する
//
// 参考分割パターン:
//
//	3人 → [3]      4人 → [4]      5人 → [5]
//	6人 → [3,3]    7人 → [3,4]    8人 → [4,4]
//	9人 → [4,5] か [3,3,3]
//	10人→ [5,5]    11人→ [4,4,3]  12人→ [4,4,4]
func Shuffle(userIDs []string) [][]string {
	if len(userIDs) < 3 {
		return nil
	}

	shuffled := make([]string, len(userIDs))
	copy(shuffled, userIDs)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// ここから下を埋めるのが宿題です ↓↓↓

	// 【ステップ1】空の groups を用意する
	// ヒント: groups := [][]string{}
	// TODO(human): groups を作る

	// 【ステップ2】for ループで 0, 4, 8, ... と進める
	// ヒント: for i := 0; i < len(shuffled); i = i + 4 {
	//          end := i + 4
	//          if end > len(shuffled) {
	//              end = len(shuffled)  // 最後で範囲を超えないように調整
	//          }
	//          groups = append(groups, shuffled[i:end])
	//        }

	// 【ステップ3】最後に return する
	// ヒント: return groups

	return nil // この行はステップ3を書いたら消す
}
