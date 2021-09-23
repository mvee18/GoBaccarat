package main

import (
	"fmt"
	"testing"
)

func TestGenerateHand(t *testing.T) {
	t.Run("testing hand generation", func(t *testing.T) {
		hand, err := GenerateHand()
		if err != nil {
			t.Error(err)
		}

		fmt.Println(hand)
	})
}

func TestDetermineValue(t *testing.T) {
	t.Run("determining value of hand", func(t *testing.T) {
		hand := []card{{
			suit:  "♠",
			value: "A",
		}, {
			suit:  "♠",
			value: "9",
		}}

		got := DetermineValue(hand)
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDetermineCoupWinner(t *testing.T) {
	t.Run("testing natural wins and natural ties", func(t *testing.T) {
		pHand := []card{{
			suit:  "♠",
			value: "A",
		}, {
			suit:  "♠",
			value: "7",
		}}

		dHand := []card{{
			suit:  "♥",
			value: "J",
		}, {
			suit:  "♥",
			value: "9",
		}}

		pT := DetermineValue(pHand)
		dT := DetermineValue(dHand)

		result := DetermineCoupWinner(&pHand, &dHand, &pT, &dT)

		if result != DealerWin {
			t.Errorf("Expected %v, got %v\n", DealerWin, result)
		}

	})
}
