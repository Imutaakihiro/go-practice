package main

import (
	"fmt"
)

// 点数を成績ランクに変換する関数（if-else if を使う）
func gradeOf(score int) string {
	if score >= 90 {
		return "S"
	} else if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else {
		return "D"
	}
}

// 成績ランクからコメントを返す関数（switch を使う）
func commentOf(grade string) string {
	switch grade {
	case "S":
		return "素晴らしい!"
	case "A":
		return "よくできました!"
	case "B":
		return "もう少し頑張ろう"
	case "C":
		return "復習しよう"
	default:
		return "補習が必要"
	}
}

func main() {
	// 名前と点数のスライス（同じ番号で対応する）
	names := []string{"田中", "佐藤", "鈴木", "山田"}
	scores := []int{85, 72, 45, 95}

	for index, name := range names {
		score := scores[index]
		grade := gradeOf(score)
		comment := commentOf(grade)
		fmt.Println("名前:", name, "点数:", score, "成績:", grade, "コメント:", comment)
	}

	// TODO(human): for range を使って names を1人ずつ回し、
	// 名前・点数・成績・コメントを1行ずつ表示する
	//
	// やること:
	//   1. for index, name := range names で1人ずつ取り出す
	//   2. score := scores[index] で対応する点数を取り出す
	//   3. grade := gradeOf(score) で成績を計算
	//   4. comment := commentOf(grade) でコメントを取得
	//   5. fmt.Println で全部まとめて表示
	//
	// ヒント（出力イメージ）:
	//   名前: 田中, 点数: 85, 成績: A, コメント: よくできました!
}
