package main

import (
	"fmt"
	"strings"
)

var REQ_END = "\r\n"

func main() {
	test_string := "GET /echo/abc HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n"

	iter := strings.SplitSeq(test_string, REQ_END)
	for partString := range iter {
		parts := strings.Split(partString, " ")

		switch strings.ToLower(parts[0]) {
		case "get":

			fmt.Println(parts[0], "0")
			fmt.Println(parts[1], "1")
			fmt.Println(parts[2], "2")
		// default:
		// 	fmt.Println(parts, "Header")
		// }

	}
	// for str := splits {
	// 	fmt.Println(str)
	//
	// }
	// fmt.Println(splits)
	// for _, str := range split {
	// 	if str == REQ_END+REQ_END {
	// 		print("Breaking")
	// 		break
	// 	}
	// 	fmt.Println(str)
	// }

	// fmt.Println(split)
}
