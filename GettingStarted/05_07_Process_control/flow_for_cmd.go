// 封包，可以自己命名(可執行程式必須使用 main)
package main

import "fmt"

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能

func main() {
	// break 跳出迴圈
	/*
		var x int
		for x = 0; x < 5; x++ {
			if x == 3 {
				break
			}

			fmt.Println(x) // 打印換行
		}
	*/

	// continue 跳過當前迴圈，進入下一次迴圈
	/*
		var x int
		for x = 0; x < 5; x++ {
			if x % 2 == 0 { // x是偶數，不印出來
				continue
			}

			fmt.Println(x)
		}
	*/

	// 實際範例：持續讓使用者輸入數字，計算總和。直到使用者輸入到 0為止
	// 無窮迴圈
	var result int = 0
	var n int
	for true {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		result += n
	}

	fmt.Println("總和：", result)
}

// 參考網址：https://www.youtube.com/watch?v=MDXGIKkDg1c&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=7
