package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const zeros = "00000"

func main() {
	id := []byte("ffykfhsq")

	password2 := []byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '}

	password := ""
	hasher := md5.New()
	var number uint64

	for found := 0; found < 8; number++ {
		hasher.Reset()
		if _, err := hasher.Write(id); err != nil {
			panic(err)
		}
		if _, err := hasher.Write([]byte(fmt.Sprintf("%d", number))); err != nil {
			panic(err)
		}

		//this could be improved by checking first two bytes of raw Sum
		sum := fmt.Sprintf("%x", hasher.Sum(nil))
		if sum[:5] == zeros {
			if len(password) < 8 {
				password = password + string(sum[5])
			}

			pos, err := strconv.Atoi(string(sum[5]))
			if err != nil || pos > 7 {
				continue
			}
			if password2[pos] == ' ' {
				password2[pos] = sum[6]
				found++
			}
		}

	}

	fmt.Println("Part1: " + password)
	fmt.Println("Part2: " + string(password2))
}
