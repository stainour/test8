package card

import "github.com/pkg/errors"

type Cards []Card

type Card struct {
	destinationCity string
	sourceCity      string
}

func NewCard(sourceCity string, destinationCity string) (*Card, error) {
	if destinationCity == "" {
		return nil, errors.New("destinationCity should be non empty!")
	}

	if sourceCity == "" {
		return nil, errors.New("sourceCity should be non empty!")
	}

	return &Card{destinationCity: destinationCity, sourceCity: sourceCity}, nil

}

func (c *Card) SourceCity() string {
	return c.sourceCity
}

func (c *Card) DestinationCity() string {
	return c.destinationCity
}
