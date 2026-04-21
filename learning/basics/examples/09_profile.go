package main

import (
	"fmt"
)

type Person struct {
	Name         string
	Age          int
	FavariteFood string
}

func main() {
	me := Person{Name: "藺牟田", Age: 22, FavariteFood: "ラーメン"}
	fmt.Println("名前;", me.Name)
	fmt.Println("年齢;", me.Age)
	fmt.Println("好きなこと:", me.FavariteFood)
}

// TODO(human): Person 型を定義して、main の中で使う
//
// 手順:
//   ① この関数の外（import の下あたり）に Person 構造体を定義する
//      type Person struct {
//          Name string
//          ...
//      }
//      - フィールドは 3〜5個、好きに決めてOK（違う型を混ぜる練習）
//      - フィールド名は大文字始まりで
//
//   ② main の中で値を作る
//      me := Person{Name: "...", Age: ..., FavoriteFood: "..."}
//
//   ③ .(ドット) でフィールドを取り出して fmt.Println で表示
//      fmt.Println("名前:", me.Name)
//      fmt.Println("年齢:", me.Age)
//
//   ④ 余裕があれば me.Age = 31 で書き換えて再表示してみる
//      (「値は後から書き換えられる」感覚を掴むため)
//
// 完成イメージ:
//   名前: ご主人様
//   年齢: 30
//   好きな食べ物: ラーメン
