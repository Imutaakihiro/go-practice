package main

import (
	"fmt"
)

func main() {
	// 商品名の一覧（スライス: 順番が大事）
	products := []string{"りんご", "バナナ", "みかん"}

	// 商品名 → 在庫数（マップ: キーで引きたい）
	stock := map[string]int{
		"りんご": 10,
		"バナナ": 5,
		"みかん": 8,
	}

	fmt.Println("=== 商品一覧 ===")
	// スライスから番号で商品名を取り出して、マップに渡して在庫数を引く
	fmt.Println("-", products[0]+":", stock[products[0]], "個")
	fmt.Println("-", products[1]+":", stock[products[1]], "個")
	fmt.Println("-", products[2]+":", stock[products[2]], "個")

	fmt.Println()
	fmt.Println("=== ぶどうを追加 ===")

	products = append(products, "ぶどう")
	stock["ぶどう"] = 3

	// 4個に増えた前提で表示
	fmt.Println("-", products[0]+":", stock[products[0]], "個")
	fmt.Println("-", products[1]+":", stock[products[1]], "個")
	fmt.Println("-", products[2]+":", stock[products[2]], "個")
	fmt.Println("-", products[3]+":", stock[products[3]], "個")
}
