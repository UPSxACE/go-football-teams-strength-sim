package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/UPSxACE/go-football-teams-strength-sim/teams"
)

func ask(reader *bufio.Reader, question string) (string, error){
	fmt.Print(question + " ")

	input, err := reader.ReadString('\n')
		
	// convert CRLF to LF
	input = strings.Replace(input, "\n", "", -1)

	return strings.Replace(input, "\r", "", -1), err
}

func main(){
	fmt.Println("Hello world")
	teams.Create("Real Madrid", 90)
	teams.Create("Barcelona", 90)
	fmt.Println(teams.List())

	
	reader := bufio.NewReader(os.Stdin)
	selectedOption := ""
	
	for selectedOption != "q" {
		input, err := ask(reader, "Type something:")

		if(err != nil){
			log.Fatalf("Unexpected error: %v", err)
		}

		selectedOption = input

		fmt.Println("You typed: ", selectedOption)
	}
}