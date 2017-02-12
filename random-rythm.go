package main

import "fmt"
import "bufio"
import "os"

import "math/rand"

func main() {
	rythmPattern := [16]string{"1", "e", "+", "a", "2", "e", "+", "a", "3", "e", "+", "a", "4", "e", "+", "a"}

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for input != "quit" {
		fmt.Println("ENTER for next pattern or 'quit' to close!")
		scanner.Scan()
		input = scanner.Text()
		if input != "quit" {
			amount := rand.Intn(len(rythmPattern))
			fmt.Println(rythmPattern)
			fmt.Println(createRandomPattern(amount))
			fmt.Println()
		} else {
			fmt.Println("Bye!")
		}
	}

}

func createRandomPattern(amount int) [16]string {
	randomPattern := [16]string{}
	for i := 0; i <= amount; i++ {
		pos := rand.Intn(len(randomPattern))
		randomPattern[pos] = "x"
	}

	for idx, value := range randomPattern {
		if value != "x" {
			randomPattern[idx] = " "
		}
	}

	return randomPattern
}
