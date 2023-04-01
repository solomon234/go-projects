package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, press enter when ready."

func main() {
	var firstNumber = Random()
	var secondNumber = Random()
	var subtraction = Random()

	StartGame(firstNumber, secondNumber, subtraction)
}

func Random() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(8) + 2
}

func StartGame(num, num2, num3 int) {
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instruction

	fmt.Println("Guess the number game \n ----------------")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take them through the games
	fmt.Println("Multiply your number by", num, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", num2, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", num3, prompt)
	reader.ReadString('\n')

	//give them the answer

	fmt.Println("the answer is", num*num2-num3)
}
