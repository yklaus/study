package loop

import "fmt"

func SampleLoop1(numbers ...int) int {
	total := 0

	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	fmt.Println("==================================")

	return total
}

func SampleLoop2(numbers ...int) int {
	total := 0

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
		total += numbers[i]
	}
	fmt.Println("==================================")

	return total
}
