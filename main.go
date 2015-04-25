package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	aikind = "AI"
	humankind = "HUMAN"
)

func playerKind(s string) (PlayerKind, error) {
	if strings.EqualFold(s, aikind) {
		return AIPlayer, nil
	}
	if strings.EqualFold(s, humankind) {
		return HumanPlayer, nil
	}
	return PlayerKind(-1), fmt.Errorf("Must be one of " + aikind + " or " + humankind)
}

func usage() string {
	return fmt.Sprintf("usage: %s [AI|HUMAN] [AI|HUMAN]", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage())
		os.Exit(2)
	}

	pk1, err := playerKind(os.Args[1])
	if err != nil {
		fmt.Println("Player 1:", err)
		fmt.Println(usage())
		os.Exit(2)
	}

	pk2, err := playerKind(os.Args[2])
	if err != nil {
		fmt.Println("Player 2:", err)
		fmt.Println(usage())
		os.Exit(2)
	}

	g := MakeGame(pk1, pk2)
	isOver := false
	winner := NoPlayer
	for !isOver {
		fmt.Print(g.String(), ("\n\n" + g.CurrentPlayer().String() + " is next to play\n"))
		g = g.Next()
		isOver, winner = g.IsOver()
	}
	fmt.Print(g.String(), "\n\nGAME OVER!\n")
	if winner != NoPlayer {
		fmt.Print(winner.String() + " wins!\n")
	} else {
		fmt.Print("The game was a draw\n")
	}
}
