package welcome

import (
	"fmt"
)

// Banner prints a banner to the console
func Banner() {
	length := 50
	height := 7
	offset := 6
	spaceBetween := 3
	msg := [][]int{}

	for i := 0; i < height; i++ {
		line := []int{}
		ci := 0
		for j := 0; j < 5; j++ {
			leter := getLetter(j, i)
			for _, l := range leter {
				fmt.Print(l)

				num := ((j) * offset) + l + ci
				fmt.Println(num)
				line = append(line, num)
			}
			ci += spaceBetween

		}
		msg = append(msg, line)
	}
	pLine(length)
	for i := 0; i < height; i++ {
		line := msg[i]
		ci := 0
		for j := 0; j < length; j++ {
			if ci < len(msg[i]) && j == line[ci] {
				fmt.Print("*")
				ci++
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	pLine(length)
}

func pLine(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func getLetter(n, i int) []int {
	switch n {
	case 0:
		return letterN(i)
	case 1:
		return letterO(i)
	case 2:
		return letterT(i)
	case 3:
		return letterE(i)
	case 4:
		return letterS(i)
	default:
		return []int{}
	}
}

func letterN(i int) []int {
	n := [][]int{
		{0, 6},
		{0, 1, 6},
		{0, 2, 6},
		{0, 3, 6},
		{0, 4, 6},
		{0, 5, 6},
		{0, 6},
	}
	return n[i]
}

func letterO(i int) []int {
	n := [][]int{
		{0, 1, 2, 3, 4, 5, 6},
		{0, 6},
		{0, 6},
		{0, 6},
		{0, 6},
		{0, 6},
		{0, 1, 2, 3, 4, 5, 6},
	}
	return n[i]
}

func letterT(i int) []int {
	n := [][]int{
		{0, 1, 2, 3, 4, 5, 6},
		{3},
		{3},
		{3},
		{3},
		{3},
		{3},
	}
	return n[i]
}

func letterE(i int) []int {
	n := [][]int{
		{0, 1, 2, 3, 4, 5, 6},
		{0},
		{0},
		{0, 1, 2, 3, 4, 5, 6},
		{0},
		{0},
		{0, 1, 2, 3, 4, 5, 6},
	}
	return n[i]
}

func letterS(i int) []int {
	n := [][]int{
		{0, 1, 2, 3, 4, 5, 6},
		{0},
		{0},
		{0, 1, 2, 3, 4, 5, 6},
		{6},
		{6},
		{0, 1, 2, 3, 4, 5, 6},
	}
	return n[i]
}
