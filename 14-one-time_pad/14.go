package main

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func main() {
	salt := []byte("ihaygndm")

	keys := make(map[int]rune)
	var foundKeys []int

	hasher := md5.New()
	for i := 0; ; i++ {
		hasher.Reset()
		hasher.Write(salt)
		hasher.Write([]byte(fmt.Sprintf("%d", i)))

		hexa := fmt.Sprintf("%x", hasher.Sum(nil))

		for j := 0; j < 2016; j++ {
			hasher.Reset()
			hasher.Write([]byte(hexa))
			hexa = fmt.Sprintf("%x", hasher.Sum(nil))
		}

		//check if has any triplet
		if tripletRunes := nTupletRunes(hexa, 3); len(tripletRunes) > 0 {
			keys[i] = tripletRunes[0]

			//if contains 5 chars, check previous entries
			if quintupletRunes := nTupletRunes(hexa, 5); len(quintupletRunes) > 0 {
				// check last 1000 entries

				for j := i - 1001; j < i; j++ {
					//check all quintuplets found (not sure if necessary, probably odds are really low, but description says nothing about it. Works both with and without this)
					for _, qRune := range quintupletRunes {
						if r, ok := keys[j]; ok && r == qRune {
							//fmt.Println(string(hexa))
							foundKeys = append(foundKeys, j)
							//fmt.Println(confirmedKeys, j, "Found key")
						}
					}
				}

				if len(foundKeys) >= 64 {
					sort.Ints(foundKeys)
					fmt.Println(foundKeys[63])
					return
				}
			}
		}
	}
}

func nTupletRunes(in string, n int) []rune {
	var res []rune
	var r rune
	count := 0
	for _, char := range in {
		if r == char {
			count++
			if count >= n {
				res = append(res, char)
				count = 1
			}
		} else {
			r = char
			count = 1
		}
	}
	return res
}
