# MASTERMIND

This is an implementation of the popular Mastermind game written in Go language.

## Game Description

The player tries to guess a hidden code that consists of a sequence of 4 numbers in the range of 1 to 6. The player has a limited number of attempts to guess the code. After each attempt, the player is informed about the number of correct numbers in correct positions and the numbers which are correct but in wrong positions.

## Installation

To compile and run this code:

1. First, make sure you have Go installed on your machine. If not, you can download and install Go from https://golang.org/dl/.

2. Copy the code into a .go file, for example, mastermind.go.

3. Open a terminal/command prompt and navigate to the directory containing the mastermind.go file.

4. To compile the program, use the command:

go build mastermind.go

5. This will produce an executable. To run the program, use the command:

./mastermind

On Windows, the command will be:

mastermind.exe

Follow the prompts in the program to play the game.

## Game Rules

The program generates a sequence of 4 numbers between 1 and 6 at the beginning of the game. The player guesses the sequence by inputting a 4 digit number, each digit between 1 and 6. For example, 1234.The player is told the number of digits that are correct and in the correct position (represented by the variable correct), and the number of digits that are correct but in the wrong position (represented by the variable incorrect). The game continues until the player guesses the sequence or exhausts all their attempts.

Remember that this game requires strategy and deduction to win. Good luck and have fun!
