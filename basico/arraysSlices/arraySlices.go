package main

import "fmt"

func main() {
	fmt.Println("arraySlices")

	array1 := [5]int{5, 4, 3, 2, 1}
	fmt.Println(array1)

	array2 := [...]int{3, 4, 5, 6, 7, 8}
	fmt.Println(array2)

	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)

	slice2 := array1[0:4]
	fmt.Println(slice2)
	array1[0] = 1001
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("------------------------------")
	slice3 := make([]int, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
