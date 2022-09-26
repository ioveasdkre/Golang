// 封包，可以自己命名(可執行程式必須使用 main)
package main

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能
import "fmt"

func main() {
	// 基本輸出練習： fmt.println(資料, 資料, ...)
	// 打印換行
	fmt.Println("Hello Golang")

	// 基本輸入練習： fmt.Scanln(&變數名稱, &變數名稱, ...)
	// &變數名稱： 取得變數的指標(Pointer)
	/*
		var x int
		// &變數名稱
		fmt.Scanln(&x)
		fmt.Println(x)
	*/

	// 整合練習： 輸入兩個數字，並輸出兩數字相加的結果
	/*
		var x int
		var y int
		fmt.Print("請輸入第一個數字：")
		fmt.Scanln(&x)
		fmt.Print("請輸入第二個數字：")
		fmt.Scanln(&y)
		var result int = x + y
		fmt.Println(result)
	*/

	var x int
	var y int
	fmt.Print("請輸入兩個數字，用空格隔開：")
	fmt.Scanln(&x, &y)
	var result int = x + y
	fmt.Println(result)
}

// 參考網址：https://www.youtube.com/watch?v=iH_fLuaGfaI&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=3
