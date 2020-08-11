package main

// GOPATH의 src 폴더를 제외한 다른 패키지를 추가할 때 상대경로 혹은 절대경로를 넣어주면 된다.
// 여러 패키지를 ()로 감싸서 한번에 import할 수 있다.
import (
	"./hello"
	"fmt"
	"strings"
)

// GOPATH의 src 폴더에 있는 패키지를 추가할 때는 패키지명만 넣어주면 된다.
//import "fmt"

func main() {
	printHello()
	variableAndConstant()

	// func test
	fmt.Println(multiply(3, 4))
	fmt.Println(lenAndUpper("klaus"))
	fmt.Println(lenAndUpper2("klaus"))

	// 반환값 변수에 저장
	// 반환값이 두 개인 경우 변수도 두개여야 한다.
	// 만약 반환값을 받고 사용하지 않으려면 _ 를 사용할 수 있다.
	length, _ := lenAndUpper("klaus")
	fmt.Println(length)

	deferTest()
}

func printHello() {
	hello.SayHello()

	// error
	//hello.sayBye()
}

func variableAndConstant() {
	// variable
	// var 변수명 타입 = 값
	// 타입 생략 가능
	// 생략시 값에 따라 자동으로 타입지정
	var hello3 string = "hello"

	// 함수 내부에서는 아래와 같이 선언할 수 있음
	hello2 := "hello2"

	// constant
	// const 상수명 타입 = 값
	const world = "world"

	// 변수를 선언해 놓고 사용하지 않으면 error가 발생한다.
	// 이럴 경우 변수명을 _로 하면 warning이 발생하지 않는다.
	var _ = "abc"

	fmt.Println(hello3, hello2, world)
}

// function
// func 함수명(arg1 타입, arg2 타입 ... ) 반환타입 {}
// argument의 타입이 같으면 타입은 한 번만 적어줘도 된다.
//func multiply(a int, b int) int {
func multiply(a, b int) int {
	return a * b
}

// function2
// 여러개의 값 반환
// func 함수명() (반환타입1, 반환타입2) {}
func lenAndUpper(name string) (int, string) {
	length, upperType := len(name), strings.ToUpper(name)
	return length, upperType
}

// function3
// naked return
// 반환타입에 변수명을 지정해주고 return을 해주면 됨
func lenAndUpper2(name string) (length int, uppercase string) {
	length, uppercase = len(name), strings.ToUpper(name)
	return
}

// function4
// 함수실행이 완료된 후 작업
// defer
func deferTest() {
	defer fmt.Println("I'm done.")
	fmt.Println("execute")
}
