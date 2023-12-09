package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
cmd:
	for {
		fmt.Print("go> ")
		if !scanner.Scan() {
			break
		}
		prompt := scanner.Text()
		switch prompt {
		case "exit":
			fmt.Println("Bye!")
			break cmd
		default:
			fmt.Println(prompt)
		}
	}
}
