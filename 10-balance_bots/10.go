package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const searchedLow, searchedHigh = 17, 61

var bots = map[int]*Bot{}
var output = map[int]int{}

func AddToOutput(i int, val int) {
	if v, ok := output[i]; ok {
		output[i] = v * val
	}
	output[i] = 1 * val
}

func GetBot(i int) *Bot {
	if b, ok := bots[i]; ok {
		return b
	}
	bots[i] = &Bot{
		ID: i,
	}
	return bots[i]
}

type Bot struct {
	ID                      int
	Low, High               *int
	LowBotOut, HighBotOut   *int
	LowContOut, HighContOut *int
}

func (b *Bot) Give(val *int) {
	if b.Low == nil {
		b.Low = val
	} else if *b.Low > *val {
		b.High = b.Low
		b.Low = val
	} else {
		b.High = val
	}
}

func (b *Bot) Propagate() {
	// check if contains all options to start propagating further
	if b.High == nil || (b.HighBotOut == nil && b.HighContOut == nil) || (b.LowBotOut == nil && b.LowContOut == nil) {
		return
	}

	// Ugly to have it here, but who cares
	if *b.Low == searchedLow && *b.High == searchedHigh {
		fmt.Println("Part 1:", b.ID)
	}

	if b.HighBotOut != nil {
		bot := GetBot(*b.HighBotOut)
		bot.Give(b.High)
		bot.Propagate()
	}

	if b.LowBotOut != nil {
		bot := GetBot(*b.LowBotOut)
		bot.Give(b.Low)
		bot.Propagate()
	}

	if b.HighContOut != nil {
		AddToOutput(*b.HighContOut, *b.High)
	}

	if b.LowContOut != nil {
		AddToOutput(*b.LowContOut, *b.Low)
	}

	b.High = nil
	b.Low = nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		fVal, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		switch parts[0] {
		case "bot": //bot 171 gives low to bot 67 and high to bot 141
			b := GetBot(fVal)

			// bot 171 gives low to bot{5} 67{6} and high to bot 141
			if lowTarget, err := strconv.Atoi(parts[6]); err != nil {
				panic(err)
			} else {
				if parts[5] == "bot" {
					b.LowBotOut = &lowTarget
				} else {
					b.LowContOut = &lowTarget
				}
			}

			// bot 171 gives low to bot 67 and high to bot{10} 141{11}
			if highTarget, err := strconv.Atoi(parts[11]); err != nil {
				panic(err)
			} else {
				if parts[10] == "bot" {
					b.HighBotOut = &highTarget
				} else {
					b.HighContOut = &highTarget
				}
			}

			b.Propagate()
		case "value": //value 67 goes to bot 17{5}
			if bID, err := strconv.Atoi(parts[5]); err != nil {
				panic(err)
			} else {
				b := GetBot(bID)
				b.Give(&fVal)
				b.Propagate()
			}
		}
	}

	fmt.Println("Part 2:", output[1]*output[0]*output[2])
}
