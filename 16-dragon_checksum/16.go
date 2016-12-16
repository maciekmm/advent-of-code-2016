package main

import (
	"bytes"
	"fmt"
)

func reverseAndFlip(s string) string {
	//unnecessary buffer creation
	var rFlipped bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			rFlipped.WriteRune('1')
		} else if s[i] == '1' {
			rFlipped.WriteRune('0')
		}
	}
	return rFlipped.String()
}

func calculateChecksum(size int, sequence string) string {
	//This is quite costly and whatnot, but will never be a bottleneck
	//This generates the sequence of concatention runes
	output := bytes.NewBufferString("0")
	for output.Len()*(len(sequence)+1) < size {
		rev := reverseAndFlip(output.String())
		output.WriteRune('0')
		output.WriteString(rev)
	}
	dragonSequence := output.String()
	//end of stupid code

	//the _ is a placeholder for dragonSequence
	fseq := sequence + "_" + reverseAndFlip(sequence) + "_"

	elements := (len(sequence)*2 + 2)

	//generate uncompressed checksum
	var checkSum []byte
	for i := 0; i < size; i += 2 {
		segmentIndex := i % elements
		fir, sec := fseq[segmentIndex], fseq[segmentIndex+1]

		if sec == '_' {
			sec = dragonSequence[0]
			dragonSequence = dragonSequence[1:]
		}

		if sec == fir {
			checkSum = append(checkSum, '1')
		} else {
			checkSum = append(checkSum, '0')
		}
	}

	//compress checksum
	for len(checkSum)%2 == 0 {
		for i := 0; i < len(checkSum); i += 2 {
			ind := (i + 1) / 2
			if checkSum[i] == checkSum[i+1] {
				checkSum[ind] = '1'
			} else {
				checkSum[ind] = '0'
			}
		}
		checkSum = checkSum[:len(checkSum)/2]
	}
	return string(checkSum)
}

func main() {
	fmt.Println("Part 1", calculateChecksum(272, "10011111011011001"))
	fmt.Println("Part 2", calculateChecksum(35651584, "10011111011011001"))
}
