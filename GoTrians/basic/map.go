package basic

import "fmt"

func lengthOfNonRepeatingSubStr(str string) int{
	maxlenth := 0
	lastoccured := make(map[byte]int)
	start := 0

	fmt.Println(lastoccured)
	for num,value := range []byte(str) {
		if lastI,ok := lastoccured[value];ok && lastI >= start{
			start = lastI + 1
		}
		if num-start+1 > maxlenth{
			maxlenth = num - start + 1
		}
		lastoccured[value] = num
	}

	return maxlenth
}

func main() {
	//var map1 = map[string]string{
	//	"name":"zhouxuan",
	//	"age":"18",
	//	"weight":"80",
	//	"telphone":"13265268546",
	//}
	////map2 := make(map[int]string)
	////var map3 map[int]string
	//
	////traverse the map
	////for key,value := range map1{
	////	fmt.Println(key,value)
	////}
	//
	////simple to validate Key exisitable
	//if username,ok := map1["nam"]; ok{
	//	fmt.Println("username is ", username)
	//}else{
	//	fmt.Println("The Key not exisited")
	//}
	//inputString := "couldyoupleaseprovide"
	//slice := make([]string,1)

	//for number := 0; number < len(inputString); number = number + 1 {
	//	value := inputString[number]
	//	for Snumber := 0; Snumber < len(slice); Snumber = Snumber + 1{
	//		Svalue := slice[Snumber]
	//		if string(value) != Svalue {
	//			slice = append(slice, string(value))
	//			fmt.Println(slice)
	//		}else{
	//			break
	//		}
	//	}
	//}
	//copy(slice,slice[len(inputString):])

	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
}
