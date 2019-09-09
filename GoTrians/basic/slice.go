package basic

import (
	"fmt"
)

func printArr(arr [7]int) {
	fmt.Println("print Array ")
	for num, value := range arr {
		fmt.Println(num, value)
	}
}

func printSlice(arr []int) {
	fmt.Println("print slices")
	for num, value := range arr {
		fmt.Println(num, value)
	}
}

func main() {
	var arr1 = [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := arr1[1:]
	s2 := s1[2:]

	printArr(arr1)
	fmt.Println("The length of arr1 and cap of it =", len(arr1), cap(arr1))
	//printSlice(s1)
	//fmt.Println("The length of s1 and cap of it =", len(s1),cap(s1))
	//printSlice(s2)
	//fmt.Println("The length of s2 and cap of it =", len(s2),cap(s2))

	s3 := append(s2, 8)
	fmt.Println("arr1 = ", arr1)
	fmt.Println("S2 = ", s2)
	fmt.Println("S3 = ", s3)

	fmt.Println("Deleting S2[1]....", s2[1])
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println("S2 is ..", s2)

}
