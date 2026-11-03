package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the program!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("How was the trip ?")
	input, _ := reader.ReadString('\n')

	fmt.Println("Rashika said :", input)
}
