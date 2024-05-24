package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var userInput = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to Fizz Buzz! Please enter an integer between 25 and 100 to serve as the ceiling of the 'Fizz Buzz' challenge to continue...")
	input, _ := userInput.ReadString('\n')
	input = strings.TrimSpace(input)

	ceiling, err := strconv.Atoi(input)
	if err != nil {
		message := fmt.Sprintf("The input string '%s' you entered was not an integer", input)
		panic(message)
	}

	if ceiling < 25 || ceiling > 100 {
		message := fmt.Sprintf("The input string '%s' you entered was not outside the desired range", input)
		panic(message)
	}

	for i := 1; i <= ceiling; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}
