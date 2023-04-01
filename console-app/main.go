package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("Menu")
	fmt.Println("-----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	fmt.Println("Press any key on the keyboard. Press q to quit.")

	char := ' '

	for unicode.ToLower(char) != 'q' {
		char, _, err = keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		if i > 0 {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}
	}

	fmt.Println("Program Exiting...")
}
