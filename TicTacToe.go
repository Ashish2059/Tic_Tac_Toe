//This is simple Tic Tac Toe game
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func play_Game() {
	val := 3
	var value_Holder1 = make([]string, val)
	var value_Holder2 = make([]string, val)
	var value_Holder3 = make([]string, val)
	for i := 0; i < 3; i++ {
		value_Holder1[i] = " "
		fmt.Printf("|")
		fmt.Print(value_Holder1[i])
		fmt.Printf("|")
	}
	fmt.Println(" ")
	for i := 0; i < 3; i++ {
		value_Holder2[i] = " "
		fmt.Printf("|")
		fmt.Print(value_Holder1[i])
		fmt.Printf("|")
	}
	fmt.Println(" ")
	for i := 0; i < 3; i++ {
		value_Holder3[i] = " "
		fmt.Printf("|")
		fmt.Print(value_Holder1[i])
		fmt.Printf("|")
	}
	fmt.Println(" ")
	count := 0
	fmt.Println("Player 1 is assigned with symbol:- O")
	fmt.Println(("Player 2 is assgned with symbol:- X"))
	fmt.Println("Enter the position value from 1-9 to fill you symbol")
	time.Sleep(3 * time.Second)
	clearScreen()
	var input int
askValue:
	count = count + 1
	if count > 9 {
		fmt.Println("--The game is Draw--")
		gameOver()
	}
	if count == 1 || count == 3 || count == 5 || count == 7 || count == 9 {
		fmt.Println("Player 1 turn:")
		fmt.Println("Enter the position to place the symbol:-")
	} else if count == 2 || count == 4 || count == 6 || count == 8 {
		fmt.Println("Player 2 turn:")
		fmt.Println("Enter the position to place the symbol:-")
	}
invalidInput:
	fmt.Scanln(&input)
	if input < 1 || input > 9 {
		fmt.Println("Enter the valid position from 1-9 ")
		fmt.Println("Enter the position again.")
		goto invalidInput
	}
	if input <= 3 {
		if value_Holder1[input-1] == "X" || value_Holder1[input-1] == "O" {
			fmt.Println("The position you have entered is already taken ")
			fmt.Println("Enter the valid position again.")
			goto invalidInput
		}
	}
	if input > 3 && input <= 6 {
		if value_Holder2[input-4] == "X" || value_Holder2[input-4] == "O" {
			fmt.Println("The position you have entered is already taken ")
			fmt.Println("Enter the valid position again.")
			goto invalidInput
		}
	}
	if input > 6 && input <= 9 {
		if value_Holder3[input-7] == "X" || value_Holder3[input-7] == "O" {
			fmt.Println("The position you have entered is already taken ")
			fmt.Println("Enter the valid position again.")
			goto invalidInput
		}
	}
	a := input
	b := count
	switch count {
	case 1, 3, 5, 7, 9:
		if a == 1 || a == 2 || a == 3 {
			value_Holder1[a-1] = "O"
		} else if a == 4 || a == 5 || a == 6 {
			value_Holder2[a-4] = "O"
		} else if a == 7 || a == 8 || a == 9 {
			value_Holder3[a-7] = "O"
		}
		clearScreen()
		fmt.Printf("You entered %d position\n", a)
		fmt.Println(value_Holder1)
		fmt.Println(value_Holder2)
		fmt.Println(value_Holder3)
		time.Sleep(time.Second)
		if value_Holder1[0] == value_Holder1[1] && value_Holder1[1] == value_Holder1[2] && value_Holder1[0] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder2[0] == value_Holder2[1] && value_Holder2[0] == value_Holder2[2] && value_Holder2[0] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder3[0] == value_Holder3[1] && value_Holder3[0] == value_Holder3[2] && value_Holder3[0] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder1[0] == value_Holder2[0] && value_Holder1[0] == value_Holder3[0] && value_Holder1[0] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder1[1] == value_Holder2[1] && value_Holder1[1] == value_Holder3[1] && value_Holder1[1] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder1[2] == value_Holder2[2] && value_Holder1[2] == value_Holder3[2] && value_Holder1[2] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder1[0] == value_Holder2[1] && value_Holder1[0] == value_Holder3[2] && value_Holder1[0] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		if value_Holder1[2] == value_Holder2[1] && value_Holder1[2] == value_Holder3[0] && value_Holder1[2] == "O" {
			fmt.Println("\n--- Player One Wins ---")
			gameOver()
		}
		goto askValue
	case 2, 4, 6, 8, 10:
		if b == 2 || b == 4 || b == 6 || b == 8 {
			if a == 1 || a == 2 || a == 3 {
				value_Holder1[a-1] = "X"
			} else if a == 4 || a == 5 || a == 6 {
				value_Holder2[a-4] = "X"
			} else if a == 7 || a == 8 || a == 9 {
				value_Holder3[a-7] = "X"
			}
		}
		if b == 10 {
			fmt.Println("--- The game is Draw ---")
			gameOver()
		}
		clearScreen()
		fmt.Printf("You entered %d position\n", a)
		fmt.Println(value_Holder1)
		fmt.Println(value_Holder2)
		fmt.Println(value_Holder3)
		time.Sleep(time.Second)
		if value_Holder1[0] == value_Holder1[1] && value_Holder1[1] == value_Holder1[2] && value_Holder1[0] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder2[0] == value_Holder2[1] && value_Holder2[0] == value_Holder2[2] && value_Holder2[0] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder3[0] == value_Holder3[1] && value_Holder3[0] == value_Holder3[2] && value_Holder3[0] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder1[0] == value_Holder2[0] && value_Holder1[0] == value_Holder3[0] && value_Holder1[0] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder1[1] == value_Holder2[1] && value_Holder1[1] == value_Holder3[1] && value_Holder1[1] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder1[2] == value_Holder2[2] && value_Holder1[2] == value_Holder3[2] && value_Holder1[2] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder1[0] == value_Holder2[1] && value_Holder1[0] == value_Holder3[2] && value_Holder1[0] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		if value_Holder1[2] == value_Holder2[1] && value_Holder1[2] == value_Holder3[0] && value_Holder1[2] == "X" {
			fmt.Println("\n--- Player Two Wins ---")
			gameOver()
		}
		goto askValue
	default:
		fmt.Println("Invalid Input. Enter the valid input ranging from 1-9")
	}
}
func gameOver() {
	fmt.Println("--------------------------")
	fmt.Println("|    G A M E  O V E R    |")
	fmt.Println("--------------------------")
	time.Sleep(5 * time.Second)
again:
	var xy string
	fmt.Println("Press c to continue or e to exit the game....")
	fmt.Scanln(&xy)
	if xy == "c" {
		play_Game()
	} else if xy == "e" {
		os.Exit(5)
	} else {
		fmt.Println("Enter the valid Input !!")
		goto again
	}
}
func main() {
A:
	clearScreen()
	fmt.Println("******************************")
	fmt.Println("*                           *")
	fmt.Println(" *                         *")
	fmt.Println("  * T A C   T A C   T O E *")
	fmt.Println(" *                         *")
	fmt.Println("*                           *")
	fmt.Println("******************************")
	fmt.Println("Press C ko continue -->>")
	var key string
	fmt.Scanln(&key)
	switch {
	case key == "c":
		clearScreen()
		fmt.Println("Welcome to TIC TAC TOE")
		play_Game()
	default:
		fmt.Println("Invalid Input")
		goto A

	}
}
