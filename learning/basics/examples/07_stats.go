package main

import (
	"fmt"
)

func minMax(nums []int) (int, int) {
	max := nums[0]
	min := nums[0]

	for _, score := range nums {
		if max < score {
			max = score
		}
		if min > score {
			min = score
		}
	}

	return max, min
}

// TODO(human): ここに minMax 関数を書く
//   シグネチャ: func minMax(nums []int) (int, int)
//   やること: スライスの最大値と最小値を求めて返す
//   ヒント:
//     - max, min := nums[0], nums[0] で最初の要素を暫定王者にする
//     - for range で2番目以降と比較し、大きければ max を更新、小さければ min を更新
//     - return max, min でカンマ区切りで2つ返す

func sumAvg(nums []int) (int, float64) {
	sum := 0
	for _, score := range nums {
		sum = sum + score
	}

	avg := float64(sum) / float64(len(nums))

	return sum, avg
}

// TODO(human): ここに sumAvg 関数を書く
//   シグネチャ: func sumAvg(nums []int) (int, float64)
//   やること: 合計と平均を求めて返す
//   ヒント:
//     - sum := 0 から始めて、for range で足していく
//     - 平均は float64(sum) / float64(len(nums)) で計算（両方 float64 に変換）
//     - return sum, avg

func main() {
	scores := []int{85, 72, 45, 95, 60}
	max, min := minMax(scores)
	sum, avg := sumAvg(scores)
	fmt.Println("点数一覧:", scores)
	fmt.Println("最高点:", max, "/", "最低点:", min)
	fmt.Println("合計:", sum, "/", "平均:", avg)

	// TODO(human): 上で作った関数を呼び出して結果を表示する
	//   1. max, min := minMax(scores) で2つ同時に受け取る
	//   2. sum, avg := sumAvg(scores) で2つ同時に受け取る
	//   3. fmt.Println で表示
	//
	// 完成イメージ:
	//   点数一覧: [85 72 45 95 60]
	//   最高点: 95 / 最低点: 45
	//   合計: 357 / 平均: 71.4
}
