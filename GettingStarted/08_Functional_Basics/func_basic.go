// 封包，可以自己命名(可執行程式必須使用 main)
package main

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能
import "fmt"

func main() {
	// 基礎函式語法練習
	/*
		test("Hello")
		test("您好")
		add(3, 4)
		add(-2, 5)
	*/

	// 計算 1 + 2 + 3 + ... + max 的函式包裝
	sum(100) // 1 + 2 + 3 + ... + 100
	sum(50)  // 1 + 2 + 3 + ... + 50
	sum(10)  // 1 + 2 + 3 + ... + 10
}

func test(mag string) {
	fmt.Println(mag)
}

func add(n1 int, n2 int) {
	var result int = n1 + n2
	fmt.Println(result)
}

func sum(max int) {
	var result int = 0
	var n int
	for n = 1; n <= max; n++ {
		result += n
	}

	fmt.Println(result)
}

// 參考網址：https://www.youtube.com/watch?v=PftKHp7r7W8&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=8
