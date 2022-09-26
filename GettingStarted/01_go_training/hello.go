// 封包，可以自己命名(可執行程式必須使用 main)
package main

// 核心函式庫，包含格式化和列印輸出或從各種 I/O 來源讀取輸入等相關的功能
import "fmt"

func main() {
	// 打印換行
	fmt.Println("Hello Golang")
}

// 撰寫程式 > 建置 (build) > 執行程式
// 終端機 → 建置： go build 程式檔案的名稱，例如：go build hello.go
// 執行程式： 輸入./加上執行檔的檔名 ，例如：./hello
// 備註：終端機需在同一層才可執行

// 參考網址：https://www.youtube.com/watch?v=yi9zp8yFULk&list=PL-g0fdC5RMbo9bdRzbKaCWYC2mXg2eEZE&index=1
