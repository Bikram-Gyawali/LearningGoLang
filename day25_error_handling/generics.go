package main

import "fmt"

func addInts(list []int) int {
	var sum int
	for _, item := range list {
		sum += item
	}
	return sum
}

func addFloats(list []float32) float32 {
	var sum float32
	for _, item := range list {
		sum += item
	}
	return sum
}

type InputType interface {
	int | float32 | float64 | uint16 | uint32 | int16 | uint64 | int32
}

// func addlist[T int | float32](list []T) T {
// 	var sum T
// 	for _, item := range list {
// 		sum += item
// 	}
// 	return sum
// }

func addlist[T InputType](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}
	return sum
}

func main() {
	fmt.Println("generics")

	fmt.Printf("sum of ints: %d\n", addInts([]int{1, 23, 4, 5, 5}))
	fmt.Printf("sum of floats: %.2f\n", addFloats([]float32{1, 23, 4232.432, 32.322}))

	fmt.Printf("sum of ints: %d\n", addlist([]int{1, 23, 14, 0, 5}))
	fmt.Printf("sum of floats: %.2f\n", addlist([]float32{1, 23, 1232.432, 32.322}))

}
