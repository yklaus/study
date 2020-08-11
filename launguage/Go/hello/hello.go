package hello

import "fmt"

// public variable
// 첫 글자가 대문자여야 다른 파일에서 사용할 수 있음...
const Hello = "A"

var Hello2 = "hello"

// public function
// 첫 글자가 대문자여야 다른 파일에서 사용할 수 있음...
func SayHello() {
	fmt.Println("Hello!!")
}

// private function
func sayBye() {
	fmt.Println("Bye!!")
}
