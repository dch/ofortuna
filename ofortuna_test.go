package ofortuna_test

import (
	"math/rand"
	"ofortuna"
	"strings"
	"testing"
	"github.com/google/go-cmp/cmp"
)

// with thanks to Carl Orff
// https://en.wikipedia.org/wiki/O_Fortuna
const oFortuna string = `O Fortuna!
%
velut luna
statu variabilis,
semper crescis
aut decrescis;
%
vita detestabilis
nunc obdurat
et tunc curat
%
ludo mentis aciem,
egestatem,
potestatem
dissolvit ut glaciem.
`

func TestGetRandomFortunes(t *testing.T) {
	t.Parallel()
	var fortunes = strings.NewReader(oFortuna)
	ft, err := ofortuna.NewFortuneTeller(fortunes)
	if err != nil {
		t.Fatal(err)
	}

	// use a known seed so we can deterministically validate our fortunes
	ft.Random = rand.New(rand.NewSource(3))

	want := "something"
	var got string = ft.GetRandomFortune()
	if want != got {
		t.Fatalf("wanted %q, got %q", want, got)
	}
}

func TestFortunesFromReader(t *testing.T) {
	t.Parallel()

	want := []string{}
	var got []string
	got, err := ofortuna.FortunesFromReader(strings.NewReader(oFortuna))

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}