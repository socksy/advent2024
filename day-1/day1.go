package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(filename string) (*os.File) {
	file, err := os.Open("input")
	check(err)
	defer file.Close()
	return(file)
}

func readLists(file io.Reader) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	var left, right []int
	for scanner.Scan() {
		nl := scanner.Text()
		nums := strings.Split(nl, "   ")
		fmt.Println("nl"+ nl)
		l, err := strconv.Atoi(nums[0])
		check(err)
		r, err := strconv.Atoi(nums[1])
		check(err)

		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func main()  {
	inputFile := openFile("input")
	left, right := readLists(inputFile)
	fmt.Printf("len=%d cap=%d\n", len(left), cap(left))
	fmt.Printf("len=%d cap=%d\n", len(right), cap(right))
}
