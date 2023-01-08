package main

import "fmt"

func main() {

	name := "world"
	// fmt는 FormatTemplate로써 출력을 위해 사용

	// 단순히 입출력을 진행
	fmt.Print("Hello, World! None NewLine")
	// Newline을 포함해서 출력을 진행
	fmt.Println("Hello, World!")
	// 문자열을 포맷에 맞게 출력한다.
	fmt.Printf("Hello, %s", name)
}
