package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLists(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	var left, right []int

	for scanner.Scan() {
		nl := scanner.Text()
		nums := strings.Split(nl, "   ")
		l, err := strconv.Atoi(nums[0])
		check(err)
		r, err := strconv.Atoi(nums[1])
		check(err)

		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func sumDifferences(left []int, right []int) uint64 {
	sort.Ints(left)
	sort.Ints(right)
	sum := uint64(0)
	for i, l := range left {
		r := right[i]

		var diff int
		if l > r {
			diff = l - r
		} else {
			diff = r - l
		}

		sum += uint64(diff)
	}
	return sum
}

func similarityScore(left []int, right []int) int {
	similarCounter := 0
	for _, l := range left {
		for _, r := range right {
			if l == r {
				similarCounter += l
			}
		}
	}
	return similarCounter
}

func main() {
	inputFile, err := os.Open("input")
	check(err)
	defer inputFile.Close()

	left, right := readLists(inputFile)
	totalDiff := sumDifferences(left, right)
	fmt.Println("total difference is", totalDiff)

	similarityScore := similarityScore(left, right)
	fmt.Println("similarity score is", similarityScore)
}
