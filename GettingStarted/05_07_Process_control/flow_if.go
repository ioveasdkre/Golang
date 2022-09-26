// 封包，可以自己命名(可執行程式必須使用 main)
package main

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能
import "fmt"

func main() {
	// 基本語法練習
	/*
		if true {
			// 打印換行
			fmt.Println("Go")
		} else {
			fmt.Println("Not Go")
		}
	*/

	// 簡易情境：ATM 檢測使用者輸入的金額是否適當
	var money int
	fmt.Println("請問想領多少錢")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Println("Too Few")
	} else if money <= 100000 {
		fmt.Println("OK")
	} else {
		fmt.Println("Too Much")
	}
	fmt.Println("執行完畢")
}

// 參考網址：https://www.youtube.com/watch?v=poQYWnS4-i0&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=5
