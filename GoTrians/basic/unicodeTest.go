package basic

import (
	"fmt"
)

func main() {
	var str string = "asdfsadfzhouxuan呀!"

	for _, value := range []byte(str) {
		fmt.Println(value)
	}
}
