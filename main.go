package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var AC AppConfig

func LoadConfig() error {
	file, err := os.Open("AppConfig")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AC)
	return err
}

func main() {
	err := LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config ! Exiting Program. ", err)
		return
	}
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
