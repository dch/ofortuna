package ofortuna

import (
	"io"
	"math/rand"
)

// with thanks to Carl Orff
// https://en.wikipedia.org/wiki/O_Fortuna
const O_Fortuna string = `O Fortuna!
velut luna
statu variabilis,
semper crescis
aut decrescis;
vita detestabilis
nunc obdurat
et tunc curat
ludo mentis aciem,
egestatem,
potestatem
dissolvit ut glaciem.
`

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