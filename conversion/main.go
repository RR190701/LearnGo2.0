package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to my pizza shop!!")
	fmt.Println("Rate our pizza from 1 to 5, please :")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
	} else {
		fmt.Printf("You rated our pizza %v out of 5.\n", rating)
	}

}
