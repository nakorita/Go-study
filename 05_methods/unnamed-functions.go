package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello")
	}() // 익명함수 호출

	func(a int, b int) {
		result := a + b
		fmt.Println(result)
	}(1, 3) // 함수 호출

	result := func(a string, b string) string {
		return a + b
	}("hello", " world!")
	fmt.Println(result)

	i, j := 10.2, 20.4
	divide := func(a float64, b float64) float64 {
		return a / b
	}(i, j)
	fmt.Println(divide)
}

// 1) 모든 함수들이 이름만 없고 그 외에 형태는 동일
// 2) 함수의 블록 마지막 브레이스(}) 뒤에 괄호를 사용해
// 		함수를 바로 호출한다.
//		이때 괄호 안에 매개변수를 넣을 수도 있다.
// 위의 익명 함수의 가장 큰 특징은
// 그 자리에서 만들고 그 자리에서 바로 실행한다.
// 익명 함수는 '기능적인 요소'만 쏙 빼와서 어디서든 가볍게
// 쓰려고 사용하는 것이다. 필요한 부분에서 어디서든 효율적으로 사용 가능.
// 선언 함수는 반환 값을 변수에 초기화하여 변수에 바로 할당 가능.
// ***익명 함수에서도 똑같은 기능을 하지만!
// 변수에 초기화한 익명 함수는 변수 이름을 함수의 이름처럼 사용 X.
// ***선언 함수와 익명 함수는 내부적으로 읽는 순서가 다른다.
// 선언 함수는 프로그램이 시작됨과 동시에 모두 읽는다.
// 하지만 익명 함수는 그 자리에서 실행되기 때문에
// 해당 함수가 실행되는 곳에서 읽는다.
// 즉, 선언 함수보다 익명 함수가 나중에 읽힌다.
