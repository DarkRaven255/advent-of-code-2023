package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day3a!")

	lines := [][]string{}

	//Read from file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//Read file rune by rune
	scanner := bufio.NewReader(file)

	line := []string{}
	for {
		if c, _, err := scanner.ReadRune(); err == nil {
			if string(c) == "\n" {
				lines = append(lines, parseToIntArray(line))
				line = []string{}
				continue
			}

			if string(c) == "." {
				line = append(line, "")
				continue
			}

			line = append(line, string(c))
			continue

		} else if err == io.EOF {
			lines = append(lines, parseToIntArray(line))
			break
		}
	}

	fmt.Printf("Result: %d", checker(lines))
}

func parseToIntArray(line []string) []string {
	for j := len(line) - 1; j > 0; j-- {

		_, e1 := strconv.Atoi(line[j-1])
		_, e2 := strconv.Atoi(line[j])

		for e1 == nil && e2 == nil {
			line[j-1] += line[j]
			break
		}
	}
	for i := 0; i < len(line)-2; i++ {
		for k := 0; k < len(line[i])-1; k++ {

			_, e1 := strconv.Atoi(line[i+k+1])
			_, e2 := strconv.Atoi(line[i+k])

			if e1 == nil && line[i+k] != "" && e2 == nil {
				line[i+k+1] = line[i+k]
			}
		}
	}

	for k := 0; k < len(line); k++ {
		if line[k] == "" {
			line[k] = "0"
		}

	}

	for k := 0; k < len(line); k++ {
		if _, err := strconv.Atoi(line[k]); err != nil && line[k] != "*" {
			line[k] = ""
		}
	}

	for k := 0; k < len(line); k++ {
		if line[k] == "0" {
			line[k] = ""
		}

	}

	return line
}
 
func checker(lines [][]string) int {
	sum := 0

	for i := 0; i < len(lines); i++ {
		shouldCount := false
		previousCuntedValue := ""
		gearRatio := 1
		partCounter := 0
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == "*" {
				//before
				if lines[i][safeCoordinates(j-1, len(lines[i]))] != "*" && lines[i][safeCoordinates(j-1, len(lines[i]))] != "" {
					if previousCuntedValue != lines[i][safeCoordinates(j-1, len(lines[i]))] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[i][safeCoordinates(j-1, len(lines[i]))])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[i][safeCoordinates(j-1, len(lines[i]))]
				}

				//after
				if lines[i][safeCoordinates(j+1, len(lines[i]))] != "*" && lines[i][safeCoordinates(j+1, len(lines[i]))] != "" {
					if previousCuntedValue != lines[i][safeCoordinates(j+1, len(lines[i]))] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[i][safeCoordinates(j+1, len(lines[i]))])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[i][safeCoordinates(j+1, len(lines[i]))]
				}

				//bottom left
				if lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j-1, len(lines[i]))] != "*" && lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j-1, len(lines[i]))] != "" {
					if previousCuntedValue != lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j-1, len(lines[i]))] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j-1, len(lines[i]))])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j-1, len(lines[i]))]
				}

				//bottom middle
				if lines[safeCoordinates(i+1, len(lines))][j] != "*" && lines[safeCoordinates(i+1, len(lines))][j] != "" {
					if previousCuntedValue != lines[safeCoordinates(i+1, len(lines))][j] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[safeCoordinates(i+1, len(lines))][j])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[safeCoordinates(i+1, len(lines))][j]
				}

				//bottom right
				if lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j+1, len(lines[i]))] != "*" && lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j+1, len(lines[i]))] != "" {
					if previousCuntedValue != lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j+1, len(lines[i]))] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j+1, len(lines[i]))])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[safeCoordinates(i+1, len(lines))][safeCoordinates(j+1, len(lines[i]))]
				}

				//top left
				if lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j-1, len(lines[i]))] != "*" && lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j-1, len(lines[i]))] != "" {
					if previousCuntedValue != lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j-1, len(lines[i]))] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j-1, len(lines[i]))])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j-1, len(lines[i]))]
				}

				//top middle
				if lines[safeCoordinates(i-1, len(lines))][j] != "*" && lines[safeCoordinates(i-1, len(lines))][j] != "" {
					if previousCuntedValue != lines[safeCoordinates(i-1, len(lines))][j] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[safeCoordinates(i-1, len(lines))][j])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)
					}
					previousCuntedValue = lines[safeCoordinates(i-1, len(lines))][j]
				}

				//top right
				if lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j+1, len(lines[i]))] != "*" && lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j+1, len(lines[i]))] != "" {
					if previousCuntedValue != lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j+1, len(lines[i]))] {
						shouldCount = true
						n, _ := strconv.Atoi(lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j+1, len(lines[i]))])
						gearRatio *= n
						partCounter++

						fmt.Printf("i: %d, j: %d, gearRatio: %d, partCounter: %d, number: %d\n", i, j, gearRatio, partCounter, n)

					}
					previousCuntedValue = lines[safeCoordinates(i-1, len(lines))][safeCoordinates(j+1, len(lines[i]))]
				}

				if shouldCount && partCounter == 2 {
					sum += gearRatio

					shouldCount = false
					fmt.Println("sum: ", sum)
				}
				gearRatio = 1
				partCounter = 0
			}
		}
	}

	return sum
}

func safeCoordinates(c int, maxLen int) int {
	if c <= 0 {
		return 0
	}
	if c >= maxLen {
		return maxLen - 1
	}
	return c
}
