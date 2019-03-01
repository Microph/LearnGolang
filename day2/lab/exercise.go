package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filterEven(nums []int) []int {
	var result []int
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
}

func printStar() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("It's not a number")
		} else {
			for i := 1; i <= num; i++ {
				for j := 0; j < i; j++ {
					fmt.Print("*")
				}
				fmt.Println()
			}
		}
	}
}

func main() {
	//Print Stars
	//printStar()

	//Filter to keep only even numbers
	nums := []int{2, 3, 4, 5, 6}
	fmt.Println(filterEven(nums))
}
