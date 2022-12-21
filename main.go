package main

// Задание 2. Сортировка пузырьком
import (
	"fmt"
)

const size = 6

func test() {
	fmt.Println("Hello men!")
}

func bubbles(in [size]int) [size]int {
	for j := size; j > 0; j-- {
		for i := 0; i < size-1; i++ {
			if in[i] > in[i+1] {
				in[i], in[i+1] = in[i+1], in[i]

			}

		}
	}
	return in
}

func main() {

	var num int
	var arr [size]int
	for i := 0; i < size; i++ {
		fmt.Printf("Input element number %d\n", i+1)
		fmt.Scan(&num)
		arr[i] = num
	}
	fmt.Println(bubbles(arr))
	test()
	var pause string
	fmt.Scan(&pause)

}
