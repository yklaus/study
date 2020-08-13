package array_slice_map_structs

import "fmt"

// array
// 변수명 := [array 크기] 타입 { 값1, 값2 ... }
// 자바, C의 array 처럼 크기를 지정해 주어야 함
func SampleArray() {
	animals := [3]string{"bear", "cat", "dog"}
	fmt.Println(animals)
	fmt.Println("==================================")

	// array 크기를 초과했을 때
	// invalid array index 4 (out of bounds for 3-element array) 발생
	// animals[4] = "abc"
}

// slice
// 변수명 := [] 타입 { 값1, 값2 ... }
// 자바, kotlin 등의 list 처럼 크기를 지정해주지 않아도 됨
func SampleSlice() {
	animals := []string{"rabbit", "hippo", "lion"}
	fmt.Println(animals)
	fmt.Println("==================================")

	// slice 크기를 초과했을 때
	// invalid array index 4 (out of bounds for 3-element array) 발생
	// compile time 잡아주지 않음
	// runtime에서 오류 발생
	//animals[4] = "abc"
}

// list 와의 차이점
// slice 에 값을 추가할 때 새로운 slice 가 return 된다.
// append 함수 사용법
func SampleSlice2() {
	animals := []string{"elephant", "tiger", "wolf"}
	fmt.Println(animals)
	fmt.Println("==================================")

	animals = append(animals, "horse")
	fmt.Println(animals)
	fmt.Println("==================================")
}

// map
// 변수 := map[key type]value_type { "키1": "값1", "키2": "값2", "키3": "값3" }
func SampleMap() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)

	// for 문으로 map의 값 출력
	for key, value := range nico {
		fmt.Println(key, value)
	}

	fmt.Println("==================================")

	// key 만 출력
	for key := range nico {
		fmt.Println(key)
	}

	fmt.Println("==================================")

	// value 만 출력
	for _, value := range nico {
		fmt.Println(value)
	}
}

// struct
/*
type struct명 struct {
	변수명1 타입1
	변수명2 타입2
	변수명3 타입3
}
*/
type person struct {
	name string
	age  int
}

func SampleStruct() {
	// 생성방법 1
	nico := person{"nico", 10}
	fmt.Println(nico)
	fmt.Println("==================================")

	// 생성방법 2
	nico2 := person{name: "nico", age: 10}
	fmt.Println(nico2)
	fmt.Println("==================================")

	// 생성방법 3
	nico3 := person{}
	nico3.name = "nico"
	fmt.Println(nico3)
	fmt.Println("==================================")
	// 값을 설정하지 않으면 빈 값 (숫자는 0) 출력
	nico3.age = 10
	fmt.Println(nico3)
	fmt.Println("==================================")
}
