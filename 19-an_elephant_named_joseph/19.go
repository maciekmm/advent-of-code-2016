package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := int64(3005290)
	//https://www.youtube.com/watch?v=uCsD3ZGzMgE
	//Josephus problem
	fmt.Println("Part 1:", input<<1+1-1<<uint(len(strconv.FormatInt(input, 2))))
	z := int64(math.Pow(float64(3), float64(len(strconv.FormatInt(input, 3))-1))) //we all love logarithms ;D
	fmt.Println("Part 2:", int64(math.Max(float64(input-z), float64(2*input-3*z))))
}
