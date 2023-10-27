package main

import "fmt"

func main7() {
	// array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr, len(arr), cap(arr))

	// slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// slice from array
	var slice2 []int = arr[1:3]
	fmt.Println(slice2, len(slice2), cap(slice2))

	// slicing
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[2:])

	// append
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice, len(slice), cap(slice))

	// append slice
	newSlice := []int{9, 10, 11}
	slice = append(slice, newSlice...)
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println("--------------------")
	fmt.Println("--------------------")

	// range
	sliceStr := []string{"a", "b", "c"}

	for i, v := range sliceStr {
		fmt.Println(i, v)
	}

	fmt.Println("--------------------")

	for _, v := range sliceStr {
		fmt.Println(v)
	}

	fmt.Println("--------------------")

	for i := range sliceStr {
		fmt.Println(i)
	}
}
