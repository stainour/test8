package card

import (
	"errors"
)

func (cards Cards) SortCards() (Cards, error) {
	if cards == nil {
		return nil, nil
	}

	if len(cards) <= 1 {
		return cards, nil
	}

	sourceCityToCard, error := toMap(cards, func(card *Card) string {
		return card.SourceCity()
	}) //N*C1

	if error != nil {
		return nil, error
	}

	destinationCityToCard, error := toMap(cards, func(card *Card) string {
		return card.DestinationCity()
	}) //N*C2

	if error != nil {
		return nil, error
	}

	var firstCard *Card = nil

	for _, card := range cards { //Max(N*C3)
		if _, ok := destinationCityToCard[card.SourceCity()]; !ok {
			firstCard = &card
			break
		}
	}

	if firstCard == nil {
		return nil, errors.New("invalid cards value")
	}

	sortedCards := make(Cards, len(cards))
	currentCard := firstCard

	for i := range cards { //N*C4
		sortedCards[i] = *currentCard
		delete(sourceCityToCard, currentCard.SourceCity())

		if len(sourceCityToCard) > 0 {
			nextCard, ok := sourceCityToCard[currentCard.DestinationCity()]

			if !ok {
				return nil, errors.New("missing slice element")
			}
			currentCard = nextCard
		}

	}
	//total N*(C1+C2+C3+C4) -> N
	return sortedCards, nil
}

func toMap(cards [] Card, getKey func(card *Card) string) (map[string]*Card, error) {
	resultingMap := map[string]*Card{}

	for i := range cards {
		card := &cards[i]
		key := getKey(card)

		if _, ok := resultingMap[key]; ok {
			return nil, errors.New("slice has duplicates")
		}
		resultingMap[key] = card
	}

	return resultingMap, nil
}
