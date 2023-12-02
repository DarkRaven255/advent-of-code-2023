package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day1a!")

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
		res, _ := getCalibrationValues(scanner.Text())
		sum += res
	}

	fmt.Println("Result: ", sum)

}

func getCalibrationValues(s string) (int, error) {

	values := []string{}

	for _, r := range s {
		if r >= '0' && r <= '9' {
			values = append(values, string(r))
		}
	}

	return strconv.Atoi(values[0] + values[len(values)-1])
}
