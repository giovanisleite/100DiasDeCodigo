package main

import (
	"fmt"
)

func main() {
	var tests int
	fmt.Scanf("%v", &tests)
	for i := 0; i < tests; i++ {
		var r1, r2 int
		fmt.Scanf("%v %v", &r1, &r2)
		fmt.Println(r1 + r2)
	}
}
