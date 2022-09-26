// 封包，可以自己命名(可執行程式必須使用 main)
package main

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能
import "fmt"

func main() {
	// 算數運算：+, -, *, /
	var x int
	x = 3*3 + 10
	fmt.Println(x) // 打印換行

	// 指定運算：=, +=, -=, *=, /=,
	x = 5
	x = x + 1 // 同 x += 1
	fmt.Println(x)

	// 單元運算：++, --
	x++
	fmt.Println(x)

	// 比較運算：>, <, >=, <=, ==, !=
	var result bool = 4 != 3
	fmt.Println(result)

	// 邏輯運算：!, &&, ||
	result = !true // 反運算
	fmt.Println("= !", result)
	result = true && false // 兩邊都成立為ture
	fmt.Println("&&", result)
	result = true || false
	fmt.Println("||", result)
}

// 參考網址：https://www.youtube.com/watch?v=cMq1aVvHZKo&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=4
