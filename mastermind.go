package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"
	colorBlue := "\033[34m"
	var number int
	var trial int
	var input [4]int
	var sequence [4]int
	stdin := bufio.NewReader(os.Stdin)
	min := 1
	max := 6

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		sequence[i] = (rand.Intn(max-min+1) + min)
	}

	//fmt.Println(sequence) //to see correct sequence

	fmt.Println(string(colorPurple), "**************Welcome to MASTERMIND!**************")
	fmt.Println(string(colorYellow), "I came up with a sequence of 4 numbers between 1 and 6. Can you guess it?")
	fmt.Println(string(colorGreen), "Enter the number of trials: ")
	fmt.Scanf("%v", &trial)
	var t int = 0
	var c int = 0
	tab := make([][]int, trial)
	for i := 0; i < len(tab); i++ {
		tab[i] = make([]int, 4)
	}
	tab2 := make([][]int, trial)
	for i := 0; i < len(tab2); i++ {
		tab2[i] = make([]int, 2)
	}
	fmt.Println(string(colorBlue), "Let's get started!Enter the first sequence: ")
	for r := 0; r < trial; {

		fmt.Fscan(stdin, &number)
		stdin.ReadString('\n')

		isInt := checkInput(number)
		if isInt == 0 {

			for i := 4; i > 0; i-- {
				input[i-1] = number % 10
				tab[t][i-1] = number % 10
				number = number / 10
			}
			t++

			correct, incorrect := check(input, sequence)
			tab2[c][0] = correct
			tab2[c][1] = incorrect
			c++
			fmt.Println(string(colorCyan), "The number of digits in correct positions: ", correct)
			fmt.Println(string(colorCyan), "The number of digits in incorrect positions: ", incorrect)
			fmt.Println(string(colorPurple), "\nRemaining trials: ", trial-r-1)

			r++
			if r == trial {
				break
			}

			time.Sleep(time.Second * 5)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			print(string(colorWhite), "Your trials: \n")
			for i := 0; i < t; i++ {
				print(string(colorWhite), "[")
				for j := 0; j < 4; j++ {
					if tab[i][j] == 1 {
						print(string(colorYellow), tab[i][j])
					}
					if tab[i][j] == 2 {
						print(string(colorBlue), tab[i][j])
					}
					if tab[i][j] == 3 {
						print(string(colorRed), tab[i][j])
					}
					if tab[i][j] == 4 {
						print(string(colorGreen), tab[i][j])
					}
					if tab[i][j] == 5 {
						print(string(colorPurple), tab[i][j])
					}
					if tab[i][j] == 6 {
						print(string(colorCyan), tab[i][j])
					}
					if tab[i][j] == 0 {
						print(string(colorWhite), tab[i][j])
					}

				}
				print(string(colorWhite), "]")
				print(string(colorWhite), " In correct position: ", tab2[i][0])
				print(string(colorWhite), " | In incorrect position: ", tab2[i][1])
				print(string(colorWhite), "\n")
			}
			fmt.Println(string(colorRed), "Enter the sequence: ")
		} else {
			fmt.Println(string(colorCyan), "Incorrect input")
			fmt.Println(string(colorWhite), "Try again (correct input is e.g. 1234 and digits between 1 to 6)")

		}

	}
	fmt.Println(string(colorRed), "You lost! The correct answear is:", sequence)

}

func check(input [4]int, v [4]int) (int, int) {
	colorBlue := "\033[34m"
	var bad int
	var good int

	for i := 0; i < 4; i++ {
		if v[i] == input[i] {
			good++
			v[i] = 0
			input[i] = 0
		}
	}

	if good == 4 {
		fmt.Println(string(colorBlue), "You win! Congratulations!!!")
		os.Exit(0)
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if input[i] == v[j] && input[i] != 0 {
				bad++
				v[j] = 0
				break
			}
		}
	}
	return good, bad
}

func checkInput(number int) int {
	var verification int
	if number > 6666 || number < 1111 {
		verification = 1
		return verification
	}
	for i := 0; i < 4; i++ {
		if number%10 == 0 || number%10 == 7 || number%10 == 8 || number%10 == 9 {
			verification = 1
			break
		}
		number = number / 10
	}
	return verification
}
