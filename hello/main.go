package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, what is your name?")
	userInput := bufio.NewReader(os.Stdin)
	name, _ := userInput.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s, welcome! I'm glad you're a fellow Gopher like me.", name)
}
