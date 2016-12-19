package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := int64(3005290)
	//https://www.youtube.com/watch?v=uCsD3ZGzMgE
	//Josephus problem
	fmt.Println(input<<1 + 1 - 1<<uint(len(strconv.FormatInt(input, 2))))
}
