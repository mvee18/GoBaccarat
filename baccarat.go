package main

import (
	"errors"
	"strconv"
)

var ErrPlayerHand = errors.New("could not generate player hand")
var ErrDealerHand = errors.New("could not generate dealer hand")

func PlayerHand() (hand []card, err error) {
	PlayerHand, err := GenerateHand()
	if err != nil {
		return PlayerHand, ErrPlayerHand
	}
	// fmt.Printf("Your hand is %v, %v\n", PlayerHand[0], PlayerHand[1])
	return PlayerHand, err
}

func DealerHand() (hand []card, err error) {
	DealerHand, err := GenerateHand()
	if err != nil {
		return DealerHand, ErrPlayerHand
	}
	//	fmt.Printf("The Dealer's Hand is {X X}, %v\n", DealerHand[1])
	// fmt.Printf("The dealer's hand is {X, X}, %v\n", DealerHand[1])
	return DealerHand, err
}

func paintConverter(s string) int {
	switch s {
	case "A":
		return 1
	case "K":
		return 10
	case "Q":
		return 10
	case "J":
		return 10
	default:
		i, _ := strconv.Atoi(s)
		return i
	}
}

const PlayerWin = 0
const Tie = 1
const DealerWin = 2

// 0 means player win, 1 means tie, 2 means dealer win.
func DetermineCoupWinner(p *[]card, d *[]card, pt *int, dt *int) int {
	if *pt >= 8 || *dt >= 8 {
		if *pt > *dt {
			return PlayerWin
		} else if *pt < *dt {
			return DealerWin
		} else {
			return Tie
		}

	} else if *pt <= 5 {

	}

	return 3
}
