package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputArr := getInputArray()
	part1Answer := getPart1Answer(inputArr)
	part2Answer := getPart2Answer(inputArr)
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func getPart2Answer(inputArr []int) int {
	defer timeTrack(time.Now(), "getPart2Answer")
	for i, val := range inputArr {
		for i2, val2 := range inputArr {
			// Don't use the same number twice
			if val+val2 > 2020 ||  i == i2 {
				continue
			}
			for i3, val3 := range inputArr {

				// Don't use the same number twice
				if i2 == i3 || i2 == i {
					continue
				}
				if val+val2+val3 == 2020 {
					return val * val2 * val3

				}
			}
		}
	}
	return 0
}


func getPart1Answer(inputArr []int) int {
	defer timeTrack(time.Now(), "getPart1Answer")
	for i, v := range inputArr {
		for i2, v2 := range inputArr {
			// Don't compare the same number to itself
			if i == i2 {
				continue
			}
			if v+v2 == 2020 {
				return v * v2
			}
		}
	}
	return 0
}

func getInputArray() []int {
	dat, err := ioutil.ReadFile("day1/data.txt")
	if err != nil {
		panic(err)
	}
	inputArr := strings.Split(string(dat), ",")
	var inputIntArr []int
	for _, num := range inputArr {
		intVal, _ := strconv.Atoi(num)
		inputIntArr = append(inputIntArr, intVal)
	}
	return inputIntArr
}
