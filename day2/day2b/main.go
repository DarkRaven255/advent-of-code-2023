package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id   int
	sets []Set
}

func main() {
	fmt.Println("Day2a!")

	sum := 0

	//Read from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += conditionChecker(decoder(scanner.Text()))
	}

	fmt.Printf("Result: %d", sum)
}

func conditionChecker(game Game) int {
	redMin := 0
	greenMin := 0
	blueMin := 0

	for _, set := range game.sets {

		if set.red > redMin {
			redMin = set.red
		}

		if set.green > greenMin {
			greenMin = set.green
		}

		if set.blue > blueMin {
			blueMin = set.blue
		}
	}

	return redMin * greenMin * blueMin
}

func decoder(s string) Game {

	game := Game{}
	str := strings.Split(s, ":")
	id, err := strconv.Atoi(strings.Trim(str[0], "Game "))
	game.id = id

	if err != nil {
		fmt.Println(err)
	}

	setsUnparsed := strings.Split(str[1], ";")

	for _, set := range setsUnparsed {
		set := strings.Split(set, ",")
		decodedSet := Set{}

		for _, singleColor := range set {
			valueAndColor := strings.Split(strings.TrimLeft(singleColor, " "), " ")

			switch valueAndColor[1] {
			case "red":
				decodedSet.red, err = strconv.Atoi(valueAndColor[0])
				if err != nil {
					fmt.Println(err)
				}
			case "green":
				decodedSet.green, err = strconv.Atoi(valueAndColor[0])
				if err != nil {
					fmt.Println(err)
				}
			case "blue":
				decodedSet.blue, err = strconv.Atoi(valueAndColor[0])
				if err != nil {
					fmt.Println(err)
				}
			}

		}
		game.sets = append(game.sets, decodedSet)
	}
	return game
}
