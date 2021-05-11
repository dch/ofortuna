package ofortuna

import (
	"io"
	"math/rand"
)

type fortuneTeller struct {
	fortune []string
	Random *rand.Rand
}

func NewFortuneTeller(reader io.Reader) (fortuneTeller, error) {
	var ft = fortuneTeller{
	}

	// get stuff from the reader, maybe fail & return error
	// return ft stuffed with goodies, shouldn't fail unless memory
	return ft, nil
}

func (fortuneTeller) GetRandomFortune() string {
	return ""
}