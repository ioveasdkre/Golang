// 封包，可以自己命名(可執行程式必須使用 main)
package main

import "fmt"

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能

func main() {
	// 基本迴圈使用
	/*
		// 無窮迴圈，Ctrl + C可停止
		for true {
			fmt.Println("Hello")
		}
	*/
	/*
		var x int = 0
		for x < 3 {
			fmt.Println(x)
			x++
		}
	*/
	/*
		var x int = 0
		for x = 0; x < 5; x += 2 {
			fmt.Println(x)
		}
	*/
	// 實際範例：計算 1 + 2 + 3 + ... + 50 的結果
	var result int = 0 // 初始值
	var x int = 1      // 起始值
	// 利用 for迴圈判斷跑到結束值
	for x <= 3 {
		fmt.Println("result:", result, "x:", x)
		result = result + x
		x++
	}
	fmt.Println(result)
}

// 參考網址：https://www.youtube.com/watch?v=u3DPQ7tdynw&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=6
