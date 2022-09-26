// 封包，可以自己命名(可執行程式必須使用 main)
package main

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能
import "fmt"

func main() {
	// 輸出訊息到終端： fmt.Println(資料, 資料, ...)
	// 打印換行
	fmt.Println(3)      // 整數 int
	fmt.Println(3.1415) // 浮點數 float64
	fmt.Println("測試中文") // 字串 string
	fmt.Println(true)   // 布林值 bool
	fmt.Println('a')    // 字符 rune
	fmt.Println("")

	var x int // 宣告變數
	x = 4
	fmt.Println(x)
	x = 10
	fmt.Println(x)
	var f float64 = 3.1415
	fmt.Println(f)
	var s string = "Hello Golang"
	fmt.Println(s)
	var b bool = true
	fmt.Println(b)
	var r rune = 'b'
	fmt.Println(r)
}

// 參考網址：https://www.youtube.com/watch?v=D6a9RpuL_UA&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=2
