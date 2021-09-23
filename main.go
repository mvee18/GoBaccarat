package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// var ErrGenerateCard = errors.New("Could not generate the card.")

type card struct {
	suit  string
	value string
}

var AllCards = make([]card, 0)

func generateCard(suit string, value string) (c card) {
	c = card{suit: suit}
	c.value = value

	_, b := SeeIfCardGenerated(AllCards, c)
	switch b {
	case true:
		c = generateCard(suitValue(), cardValue())
		return c
	}
	AllCards = append(AllCards, c)
	return c
}

func cardValue() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	value := r1.Intn(13-1) + 1
	f, err := convertValue(value)
	if err != nil {
		log.Fatal("Could not convert the value.")
	}
	return f
}

func convertValue(i int) (paint string, err error) {
	switch i {
	case 1:
		paint := "A"
		return paint, nil
	case 11:
		paint := "J"
		return paint, nil
	case 12:
		paint := "Q"
		return paint, nil
	case 13:
		paint := "K"
		return paint, nil
	default:
		paint := strconv.Itoa(i)
		return paint, nil
	}
}

func suitValue() (s string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	suit := r1.Intn(3)
	switch suit {
	case 0:
		s := "♠"
		return s
	case 1:
		s := "♥"
		return s
	case 2:
		s := "♦"
		return s
	case 3:
		s := "♣"
		return s
	default:
		s := "0"
		log.Fatal("Did not get a suit.")
		return s
	}
}

func GenerateHand() (hand []card, err error) {
	hand = make([]card, 2)
	hand[0] = generateCard(suitValue(), cardValue())
	time.Sleep(100 * time.Millisecond)
	hand[1] = generateCard(suitValue(), cardValue())
	return hand, err
}

func SeeIfCardGenerated(s []card, c card) (int, bool) {
	for i, item := range s {
		if item == c {
			return i, true
		}
	}
	return -1, false
}

func DetermineValue(hand []card) (total int) {
	for _, c := range hand {
		val := paintConverter(c.value)
		total += val
	}
	return total % 10
}

func main() {
	d, err := DealerHand()
	if err != nil {
		panic(err)
	}

	p, err := PlayerHand()
	if err != nil {
		panic(err)
	}

	dTotal := DetermineValue(d)
	pTotal := DetermineValue(p)

	fmt.Printf("The player hand is %v, with a total of %d\n", p, pTotal)
	fmt.Printf("The player hand is %v, with a total of %d\n", d, dTotal)
}
