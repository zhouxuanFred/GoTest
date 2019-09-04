package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
)

// switch practices
func swithTest(a,b int, op string) (result int, err error) {
	switch op {
	case "+":
		return a+b, nil
	case "-":
		return a-b, nil
	case "*":
		return a*b, nil
	case "/":
		return a/b, nil
	default:
		return 0, fmt.Errorf("unsupport operation: %s", op)
	}
}

// body-function practices
func apply(test2 func(int, int) int, a, b int) int {
	//get the fucntion name
	i := reflect.ValueOf(test2).Pointer()
	opName := runtime.FuncForPC(i).Name()

	fmt.Printf("The function name is %s with args (%d, %d) \n", opName, a,b)
	return test2(a,b)
}

//calculate "+"
func Plus(a, b int) int {
	return a + b
}

// calculate
func pow(a ,b int ) int{
	return int(math.Pow(float64(a), float64(b)))
}

// a,b = b, a
func swap (a,b int) (int, int) {
	return b,a
}

func binaryMethod(input int) string{
	result := ""

	for ;input>0; input /= 2{
		lsb := input % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	 if result, err := swithTest(1,2, "9"); err != nil {
	 	fmt.Println("Error info", err)
	 }else {
	 	fmt.Printf("The result is %d", result)
	 }

	 fmt.Printf("The operation's result is %d\n", apply(Plus,4,5))
	 fmt.Printf("The operation's result is %d\n", apply(pow,4,5))

	 a,b := 3,4
	 a,b = swap(a,b)
	 fmt.Printf("a is %d, b is %d\n", a,b)

	// var binaryOutput, input int
	// fmt.Printf("The binary of %d is %s\n",input, binaryOutput )

	 fmt.Printf("The binary is %s",binaryMethod(10))

}
