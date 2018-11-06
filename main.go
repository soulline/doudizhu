package main

import (
	"doudizhu/computer"
	"fmt"
)

func main() {
	initValues := computer.CreateNew()
	computer.Shuffle(initValues)
	dispacherV := computer.Dispacther(2, initValues)
	fmt.Println(dispacherV)
	cardsA := []string{"A3", "B3", "C3", "A4", "B4", "C4", "A5", "B5", "A5", "A6", "B6", "A6", "A11", "A7", "B12", "B7"}
	ashowMode := computer.ParseCardsInSize(cardsA)
	cardsB := []string{"A4", "B4", "C4", "A5", "B5", "C5", "A6", "B6", "A6", "A7", "B7", "A7", "A11", "A10", "B12", "13"}
	bshowMode := computer.ParseCardsInSize(cardsB)
	fmt.Println("\nA玩家:", ashowMode.CompareValue)
	fmt.Println("B玩家:", bshowMode.CompareValue)
}
