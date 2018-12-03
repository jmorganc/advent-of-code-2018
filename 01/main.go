package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	checkError(err)
	defer f.Close()

	var i int
	hasDupe := map[int]bool{0: false}
	s := bufio.NewScanner(f)
	for {
		for s.Scan() {
			intVal, err := strconv.Atoi(s.Text())
			checkError(err)
			fmt.Println(intVal)
			i += intVal

			if _, ok := hasDupe[i]; ok {
				fmt.Println("Dupe found: ", i)
				os.Exit(0)
			} else {
				hasDupe[i] = false
			}
		}
		fmt.Println("Final frequency:", i)
		f.Seek(0, 0)
		s = bufio.NewScanner(f)
	}
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
