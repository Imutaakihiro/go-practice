package main

import (
	"fmt"
)

func main() {
	// ゼロ値の確認: 宣言だけしてみる
	var emptyHeight float64
	var emptyWeight int
	fmt.Println("ゼロ値の確認 — 身長:", emptyHeight, "体重:", emptyWeight)

	// 身長は小数(float64)、体重は整数(int)で持つ
	var height float64 = 1.7
	var weight int = 65
	bmi := float64(weight) / (height * height)

	fmt.Println("身長(m):", height)
	fmt.Println("体重(kg):", weight)
	fmt.Println("BMI:", bmi)
}
