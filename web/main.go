package main

import (
	"fmt"
	"io"

	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Welcome to the web!")
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error occurred:", err)
		panic(err)
	}
	//closing the connection
	fmt.Printf("Response is of type %T\n:", res)
	defer res.Body.Close()
	fmt.Println("Response received successfully!")
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error occurred while reading response body:", err)
		panic(err)
	}
	fmt.Println("Response data:", string(data))
}
