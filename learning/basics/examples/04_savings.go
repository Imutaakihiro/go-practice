package main

import (
	"fmt"
)

// ポインターを使わない版（コピーが渡るので元は変わらない）
func depositCopy(balance int, amount int) {
	balance = balance + amount
}

// ポインターを使う版（住所を渡すので元の値が変わる）
func depositPtr(balancePtr *int, amount int) {
	*balancePtr = *balancePtr + amount
}

func main() {
	fmt.Println("=== ポインターなし版 ===")
	balance1 := 1000
	fmt.Println("入金前:", balance1)
	depositCopy(balance1, 500)
	fmt.Println("入金後:", balance1)

	fmt.Println()

	fmt.Println("=== ポインター版 ===")
	balance2 := 1000
	fmt.Println("入金前:", balance2)
	depositPtr(&balance2, 500) // & で住所を渡す
	fmt.Println("入金後:", balance2)
}
