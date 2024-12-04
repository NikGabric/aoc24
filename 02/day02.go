package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/NikGabric/aoc24/helpers"
)

func Task01() {
	f, err := os.ReadFile("./02/in.txt")
	helpers.HandleErr(err)

	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	safeCount := 0
	for scanner.Scan() {
		strVals := strings.Fields(scanner.Text())
		var vals []int
		for _, val := range strVals {
			v, _ := strconv.Atoi(val)
			vals = append(vals, v)
		}

		if isSafe(vals) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func Task02() {
	f, err := os.ReadFile("./02/in.txt")
	helpers.HandleErr(err)

	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	safeCount := 0
	for scanner.Scan() {
		strVals := strings.Fields(scanner.Text())
		var vals []int
		for _, val := range strVals {
			v, _ := strconv.Atoi(val)
			vals = append(vals, v)
		}

		if isSafe(vals) {
			safeCount++
		} else {
			for i := 0; i < len(vals); i++ {
				newVals := helpers.Remove(vals, i)
				safe := isSafe(newVals)
				if safe {
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println(safeCount)
}

func isSafe(vals []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(vals); i++ {
		diff := vals[i] - vals[i-1]
		if helpers.Abs(diff) < 1 || helpers.Abs(diff) > 3 {
			return false
		}

		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}

	return decreasing || increasing
}