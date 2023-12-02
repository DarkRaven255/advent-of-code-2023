package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day1b!")

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

func replaceWordsWithNumbers(s string) string {

	str := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, v := range str {
		index := strings.Index(s, v)

		if index != -1 {
			s = strings.Replace(s, v, fmt.Sprintf("%s%s%s", string(v[0]), strconv.Itoa(i+1), string(v[len(v)-1])), -1)
		}
	}

	return s
}

func getCalibrationValues(s string) (int, error) {

	s = replaceWordsWithNumbers(s)

	values := []string{}

	for _, r := range s {
		if r >= '0' && r <= '9' {
			values = append(values, string(r))
		}
	}

	return strconv.Atoi(values[0] + values[len(values)-1])
}
