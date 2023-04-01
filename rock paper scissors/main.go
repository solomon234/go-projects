package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(2)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ = reader.ReadString('\n')

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	} else if playerChoice == "paper" {
		playerValue = PAPER
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer choose Rock")
		break
	case PAPER:
		fmt.Println("Computer choose Paper")
		break
	case SCISSORS:
		fmt.Println("Computer choose Scissors")
		break
	}
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
