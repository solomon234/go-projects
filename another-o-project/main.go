package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		if ClearInputs(userInput) == "quit" {
			break
		}
		fmt.Println(doctor.Response(userInput))
	}

}

func ClearInputs(str string) string {
	return strings.Replace(strings.Replace(str, "\n", "", -1), "\r\n", "", -1)
}
