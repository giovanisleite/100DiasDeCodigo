package main

import (
    "math/big"
	"fmt"
	"strings"
)

func main() {
	for true {
		var d, n string
		fmt.Scanf("%s %s", &d, &n)
		if d == n && n == "0" {
			return
		}
        answer_str := strings.Replace(n, d, "", -1)
        if answer_str == "" {
            answer_str = "0"
        }
        answer := new(big.Int)
		answer, _ = answer.SetString(answer_str, 10)
		fmt.Println(answer)
	}

}
