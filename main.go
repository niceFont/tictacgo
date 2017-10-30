package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board          [][]string
	over           bool
	previousplayer string
	currentplayer  string
	moveCount      int
}

var game *Game

func init() {
	game = new(Game)
	game.over = false
	game.previousplayer = "O"
	game.currentplayer = "X"
	game.moveCount = 0
	game.board = [][]string{
		{"#", "#", "#"},
		{"#", "#", "#"},
		{"#", "#", "#"},
	}

}

func (g *Game) checkIfOver() bool {
	for i := 0; i < len(g.board); i++ {
        if full := checkFields(g.board[i][0], g.board[i][1], g.board[i][2]); full {
			return true
		}
        if full := checkFields(g.board[0][i], g.board[1][i], g.board[2][i]); full {
			return true
		}
	}

    if full := checkFields(g.board[0][0], g.board[1][1], g.board[2][2]); full {
		return true
	}

    if full := checkFields(g.board[0][2], g.board[1][1], g.board[2][0]); full {
		return true
	}

	if g.moveCount >= 9 {
		return true
	}
	return false
}

func checkFields(field1, field2, field3 string) bool {
    if field1 != "#" && field1 == field2 && field2 == field3 {
        return true
    } else {
        return false
    }
}

func (g *Game) move(playmove int) error {

	var fieldtaken error = errors.New("Field is already taken pick a new one: ")

	if playmove > 9 || playmove < 1 {
		return errors.New("Invalid input please enter a valid Number from 1-9: ")
	}

	if playmove <= 3 {
		if g.board[0][playmove-1] == "#" {
			g.board[0][playmove-1] = g.currentplayer
			g.moveCount += 1
			return nil
		}
		return fieldtaken
	}
	if playmove > 3 && playmove <= 6 {
		if g.board[1][playmove-4] == "#" {
			g.board[1][playmove-4] = g.currentplayer
			g.moveCount += 1
			return nil
		}
		return fieldtaken
	}
	if playmove > 6 {
		if g.board[2][playmove-7] == "#" {
			g.board[2][playmove-7] = g.currentplayer
			g.moveCount += 1
			return nil
		}
		return fieldtaken
	}
	g.moveCount += 1
	return nil
}

func (g *Game) printBoard() {
	for i := 0; i < len(g.board); i++ {
		for j := 0; j < len(g.board[i]); j++ {
			fmt.Printf("%s", g.board[i][j])
			if j != 2 {
				fmt.Print("|")
			} else {
				fmt.Print("\n")
			}
		}
		if i != 2 {
			fmt.Println("=====")
		}
	}
}

func (g *Game) switchplayers() {
	g.currentplayer, g.previousplayer = g.previousplayer, g.currentplayer
}

func main() {
	fmt.Println(
		`Welcome to basic TicTacToe!
		 The Board-Fields are indexed from 1-9.
		 To select a Field just enter the number,
		 and wait for your turn.
	`)
	game.printBoard()

	for !game.over {

		fmt.Printf("Select your Field: %s\n", game.currentplayer)

		var input int

		fmt.Scanf("%d", &input)
		err := game.move(input)
		if err != nil {
			fmt.Println(err)
		} else {
			game.printBoard()
			fmt.Println(game.moveCount)
			if ok := game.checkIfOver(); ok {
				game.over = true
				if game.moveCount != 9 {
					fmt.Printf("Game Over, %s Won!\n", game.currentplayer)
				} else {
					fmt.Printf("Game Over, Its a Draw")
				}
			} else {
				game.switchplayers()
			}
		}
	}
	fmt.Scan()
}
