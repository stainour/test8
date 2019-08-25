package main

import (
	"github.com/stainour/test8/card"
	"reflect"
	"strconv"
	"testing"
)

var testCards card.Cards

const cardCount = 10000

func init() {
	testCards = make(card.Cards, cardCount)

	for i := 1; i <= cardCount; i++ {
		testCards[i-1] = newCard(strconv.Itoa(i-1), strconv.Itoa(i))
	}
}

func TestCardsHasDublicates(t *testing.T) {
	cards := append(testCards, testCards...)
	_, error := cards.SortCards()

	if error == nil {
		t.Error("error is nil")
	}

	if error.Error() == "" {
		t.Error("error is empty")
	}
}

func TestCardsMissingElement(t *testing.T) {

	badCards := make(card.Cards, cardCount)
	copy(badCards, testCards)

	badCards = append(badCards[:cardCount/3], badCards[2*cardCount/3:]...)

	_, error := badCards.SortCards()
	if error == nil {
		t.Error("error is nil")
	}

	if error.Error() == "" {
		t.Error("error is empty")
	}
}

func TestShuffledCards(t *testing.T) {
	shuffledCards := make(card.Cards, cardCount)
	copy(shuffledCards, testCards)

	for i := 0; i < cardCount-2; i += 2 {
		shuffledCards[i], shuffledCards[i+1] = shuffledCards[i+1], shuffledCards[i]
	}
	sortedCards, _ := shuffledCards.SortCards()
	validateCards(sortedCards, testCards, t)
}

func TestSingleCard(t *testing.T) {
	cards := card.Cards{newCard("1", "2")}
	sortedCards, _ := cards.SortCards()
	validateCards(sortedCards, cards, t)
}

func TestNilCards(t *testing.T) {
	var cards card.Cards = nil
	sortedCards, _ := cards.SortCards()
	if sortedCards != nil {
		t.Error("sortedCards slice is not nil")
	}
}

func TestAlreadySortedCards(t *testing.T) {
	sortedCards, _ := testCards.SortCards()
	validateCards(sortedCards, sortedCards, t)
}

func TestReverseCards(t *testing.T) {
	reverse := make(card.Cards, cardCount)

	for i := range testCards {
		reverse[i] = testCards[cardCount-1-i]
	}

	sortedCards, _ := reverse.SortCards()
	validateCards(sortedCards, testCards, t)
}

func TestEmptyCards(t *testing.T) {
	sortedCards, _ := make(card.Cards, 0).SortCards()
	validateCards(sortedCards, make(card.Cards, 0), t)
}

func newCard(sourceCity string, destinationCity string) card.Card {
	newCard, _ := card.NewCard(sourceCity, destinationCity)
	return *newCard
}

func validateCards(sortedCards card.Cards, compareToCards card.Cards, t *testing.T) {
	if !reflect.DeepEqual(sortedCards, compareToCards) {
		t.Error("sortedCards slice is not the same as compareToCards")
	}
}
