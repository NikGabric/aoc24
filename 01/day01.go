package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Task01() {
	f, err := os.ReadFile("./01/in.txt")
	handleErr(err)

	var list1, list2 []int
	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		v1, _ := strconv.Atoi(values[0])
		v2, _ := strconv.Atoi(values[1])

		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}

	sum := 0
	l := len(list1)
	for i := 0; i < l; i++ {
		min1 := slices.Min(list1)
		min2 := slices.Min(list2)
		sum += abs(min1 - min2)
		list1 = remove(list1, indexOf(list1, min1))
		list2 = remove(list2, indexOf(list2, min2))
	}

	fmt.Println(sum)
}

func Task02() {
	f, err := os.ReadFile("./01/in.txt")
	handleErr(err)

	var list1, list2 []int
	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())
		v1, _ := strconv.Atoi(values[0])
		v2, _ := strconv.Atoi(values[1])

		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}

	score := 0
	for i := 0; i < len(list1); i++ {
		leftItem := list1[i]
		count := 0
		for j := 0; j < len(list2); j++ {
			if leftItem == list2[j] {
				count += 1
			}
		}
		score += leftItem * count
	}

	fmt.Println(score)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func indexOf(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
