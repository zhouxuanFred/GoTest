package basic

import "fmt"

func arrayTest(arr *[5]int){
	arr[3] = 100
	for count, value := range arr{
		fmt.Println(count,value)
	}
}

func main() {
	arr1 := [5]int{9,1,2,3,4}
	var maxi,maxnum int = 1,1

	arrayTest(&arr1)
	for i,v := range arr1 {
		 if v > maxnum {
		 	maxi, maxnum = i,v
		 }
	}
	fmt.Println(maxi, maxnum)

	var arr2 = [...]int{1,2,3,4,5,6,7,8,9}
	var s = arr2[2:4]
	fmt.Println("The slices = ", s)
}
