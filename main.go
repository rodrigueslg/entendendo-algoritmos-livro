package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 1}
	sum := sumArray(arr, 0)
	msg := fmt.Sprintf("The sum of %v is %d:", arr, int(sum))
	fmt.Println(msg)
}

func sumArray(numbers []int, fromIndex int) int {
	if len(numbers) == 0 || len(numbers) == fromIndex {
		return 0
	}
	return numbers[fromIndex] + sumArray(numbers, fromIndex+1)
}
