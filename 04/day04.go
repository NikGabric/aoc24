package day04

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func Task01() {
	f, _ := os.ReadFile("./04/in.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(f)))

	var matrix [][]string
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "")
		matrix = append(matrix, temp)
	}

	count := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			// ->
			if x <= len(matrix[y])-4 {
				word := strings.Join(matrix[y][x:x+4], "")
				count += isWord(word, "XMAS")
			}

			// <-
			if x >= 4 {
				temp := make([]string, 4)
				copy(temp, matrix[y][x-4:x])
				slices.Reverse(temp)
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}

			// ↓
			if y <= len(matrix)-4 {
				temp := make([]string, 4)
				for k := range 4 {
					temp = append(temp, matrix[y+k][x])
				}
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}

			// ↑
			if y >= 3 {
				temp := make([]string, 4)
				for k := range 4 {
					temp = append(temp, matrix[y-k][x])
				}
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}

			// ↘
			if x <= len(matrix[y])-4 && y <= len(matrix)-4 {
				temp := make([]string, 4)
				for k := range 4 {
					temp = append(temp, matrix[y+k][x+k])
				}
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}

			// ↙
			if x >= 3 && y <= len(matrix)-4 {
				temp := make([]string, 4)
				for k := range 4 {
					temp = append(temp, matrix[y+k][x-k])
				}
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}

			// ↖
			if x >= 3 && y >= 3 {
				temp := make([]string, 4)
				for k := range 4 {
					temp = append(temp, matrix[y-k][x-k])
				}
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}

			// ↗
			if x <= len(matrix[y])-4 && y >= 3 {
				temp := make([]string, 4)
				for k := range 4 {
					temp = append(temp, matrix[y-k][x+k])
				}
				word := strings.Join(temp, "")
				count += isWord(word, "XMAS")
			}
		}
	}

	fmt.Println(count)
}

type Coords struct {
	x int
	y int
}

type WordCoords struct {
	start Coords
	end   Coords
}

func Task02() {
	f, _ := os.ReadFile("./04/in.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(f)))

	var matrix [][]string
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), "")
		matrix = append(matrix, temp)
	}

	count := 0

	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[y])-1; x++ {
			if matrix[y][x] == "A" {
				count += checkMas([]string{
					matrix[y-1][x-1],
					matrix[y-1][x+1],
					matrix[y+1][x+1],
					matrix[y+1][x-1],
				})
			}
		}
	}

	fmt.Println(count)
}

func checkMas(x []string) int {
	diag1 := (x[0] == "M" && x[2] == "S") || (x[2] == "M" && x[0] == "S")
	diag2 := (x[1] == "M" && x[3] == "S") || (x[3] == "M" && x[1] == "S")
	if diag1 && diag2 {
		return 1
	}
	return 0
}

func isWord(w string, word string) int {
	if w == word {
		return 1
	}
	return 0
}
