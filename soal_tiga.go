package main

import (
	"fmt"
	"strings"
)

var initStr string

func main() {
	//var _jm int
	var uString string

	initStr = "Berkesinambungan"

	uString = strings.ToLower(initStr)

	order := make([]byte, 0, 256)
	for _, dt := range uString {
		order = append(order, byte(dt))
	}

	res := make(map[string]int)
	for _,_item := range order {
		_,exist := res[strings.ToUpper(string(_item))]
		if exist {
			res[strings.ToUpper(string(_item))] += 1
		} else {
			res[strings.ToUpper(string(_item))] = 1	
		}
	}
	fmt.Printf("Input String: %s\n",initStr)
	fmt.Println("Duplikat")
	for _key,_val := range res {
		if _val > 1{
			fmt.Printf("%s: %d\n", _key, _val)
		}
	}
}
