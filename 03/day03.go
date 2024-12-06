package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Tasks() {
	f, _ := os.ReadFile("./03/in.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	var text string
	for scanner.Scan() {
		text += string(scanner.Text())
	}

	sol1 := solve1(text)
	sol2 := solve2(text)

	fmt.Println(sol1)
	fmt.Println(sol2)
}

func solve1(text string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(text, -1)
	var sum int

	for _, mul := range matches {
		s := strings.Replace(strings.Replace(mul, "mul(", "", 1), ")", "", 1)
		v := strings.Split(s, ",")
		v1, _ := strconv.Atoi(v[0])
		v2, _ := strconv.Atoi(v[1])

		sum += v1 * v2
	}

	return sum
}

func solve2(text string) int {
	strings.ReplaceAll(text, "\n", " ")
	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don\'t\(\))`)
	subs := re.FindAllStringSubmatch(text, -1)
	var sum int

	active := true
	for _, pat := range subs {
		if strings.Contains(pat[0], "don't") {
			active = false
			continue
		} else if strings.Contains(pat[0], "do") {
			active = true
			continue
		} else if active {
			s := strings.Replace(strings.Replace(pat[0], "mul(", "", 1), ")", "", 1)
			v := strings.Split(s, ",")
			v1, _ := strconv.Atoi(v[0])
			v2, _ := strconv.Atoi(v[1])

			sum += v1 * v2
		}
	}

	return sum
}
