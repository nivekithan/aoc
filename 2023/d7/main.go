package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Round struct {
	hand string
	bid  int
}

type RoundValue []int

func readInput() []Round {
	fileName := "d7.data"

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	allHands := []Round{}

	for scanner.Scan() {
		line := scanner.Text()

		splittedLine := strings.Split(line, " ")

		hand := splittedLine[0]
		bidInStr := splittedLine[1]

		bid, err := strconv.Atoi(bidInStr)

		if err != nil {
			log.Fatal(err)
		}

		allHands = append(allHands, Round{hand: hand, bid: bid})
	}

	return allHands
}

func part1() {
	allHands := readInput()

	slices.SortStableFunc(allHands, compareRounds)

	sum := 0
	for i, round := range allHands {
		sum += (i + 1) * round.bid
	}

	fmt.Println(sum)
}

func compareRounds(left Round, right Round) int {
	leftHandValue := getRoundValue(&left)
	rightHandValue := getRoundValue(&right)

	isSameTypeOfHand := leftHandValue[0] == rightHandValue[0]

	if !isSameTypeOfHand {
		return leftHandValue[0] - rightHandValue[0]
	}

	for i := 0; i < len(leftHandValue); i++ {
		leftValue := leftHandValue[i]
		rightValue := rightHandValue[i]

		isSame := leftValue == rightValue

		if !isSame {
			return leftValue - rightValue
		}
	}

	log.Fatal("Unreachable")
	return 0
}

const (
	Two = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	HighCard = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var RuneToCardValue = map[rune]int{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': Ten,
	'J': Jack,
	'Q': Queen,
	'K': King,
	'A': Ace,
}

type CardOccurances struct {
	cardValue  int
	occurances int
}

func getRoundValue(round *Round) []int {
	hand := round.hand

	cardOccurances := map[int]int{}
	for _, r := range hand {
		cardValue := RuneToCardValue[r]

		if cardValue == 0 {
			log.Fatalf("Unknown rune %v", string(r))
		}

		cardOccurances[cardValue] += 1
	}

	allCardOccurances := []CardOccurances{}

	for cardValue, occ := range cardOccurances {
		allCardOccurances = append(allCardOccurances, CardOccurances{cardValue: cardValue, occurances: occ})
	}

	slices.SortFunc(allCardOccurances, func(a CardOccurances, b CardOccurances) int {
		isOccuranceSame := a.occurances-b.occurances == 0

		if isOccuranceSame {
			return -(a.cardValue - b.cardValue)
		}

		return -(a.occurances - b.occurances)
	})

	cardValue := []int{}

	if isFiveOfAKind(&allCardOccurances) {
		cardValue = append(cardValue, FiveOfAKind)
	} else if isFourOfAKind(&allCardOccurances) {
		cardValue = append(cardValue, FourOfAKind)
	} else if isFullHouse(&allCardOccurances) {
		cardValue = append(cardValue, FullHouse)
	} else if isThreeOfKind(&allCardOccurances) {
		cardValue = append(cardValue, ThreeOfAKind)
	} else if isTwoPair(&allCardOccurances) {
		cardValue = append(cardValue, TwoPair)
	} else if isOnePair(&allCardOccurances) {
		cardValue = append(cardValue, OnePair)
	} else if isHighCard(&allCardOccurances) {
		cardValue = append(cardValue, HighCard)
	}

	for _, r := range hand {
		runeValue := RuneToCardValue[r]
		cardValue = append(cardValue, runeValue)
	}

	return cardValue
}

func isFiveOfAKind(allCardOccurances *[]CardOccurances) bool {

	if (*allCardOccurances)[0].occurances == 5 {
		return true
	}

	return false
}

func isFourOfAKind(allCardOccurances *[]CardOccurances) bool {
	if (*allCardOccurances)[0].occurances == 4 {
		return true
	}

	return false
}

func isFullHouse(allCardOccurances *[]CardOccurances) bool {
	if (*allCardOccurances)[0].occurances == 3 && (*allCardOccurances)[1].occurances == 2 {
		return true
	}

	return false
}

func isThreeOfKind(allCardOccurances *[]CardOccurances) bool {
	if (*allCardOccurances)[0].occurances == 3 {
		return true
	}

	return false
}

func isTwoPair(allCardOccurances *[]CardOccurances) bool {
	if (*allCardOccurances)[0].occurances == 2 && (*allCardOccurances)[1].occurances == 2 {
		return true
	}

	return false
}

func isOnePair(allCardOccurance *[]CardOccurances) bool {
	if (*allCardOccurance)[0].occurances == 2 {
		return true
	}

	return false
}

func isHighCard(allCardOccurance *[]CardOccurances) bool {
	return true
}

func main() {
	part1()
}
