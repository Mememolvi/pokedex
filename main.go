package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commandMap := getCommandMap()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("pokedex >")
		input, _ := reader.ReadString('\n') // read command
		input = strings.Trim(input, "\n")   // trim delimiter
		// fmt.Print(input)
		// do what commadn says
		inputArgs := strings.Split(input, " ")
		v, ok := commandMap[inputArgs[0]]
		if ok {
			var err error
			if len(inputArgs) > 1 {
				err = v.callback(inputArgs[1:])
			} else {
				err = v.callback(nil)
			}
			if err != nil {
				fmt.Println("Failed To Execute command please try again, with error :", err)
			}
		} else {
			fmt.Println("Please enter a valid command!")
		}

	}
}
