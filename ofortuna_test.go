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
	fortunes := strings.NewReader(oFortuna)
	ft, err := ofortuna.NewFortuneTeller(fortunes)
	if err != nil {
		t.Fatal(err)
	}

	// use a known seed so we can deterministically validate our fortunes
	ft.Random = rand.New(rand.NewSource(4))

	want := "velut luna\nstatu variabilis,\nsemper crescis\naut decrescis;\n"
	var got string = ft.GetRandomFortune()
	if want != got {
		t.Fatalf("wanted %q, got %q", want, got)
	}
}

func TestFortunesFromReader(t *testing.T) {
	t.Parallel()
	want := []string{
		"O Fortuna!\n",
		"velut luna\nstatu variabilis,\nsemper crescis\naut decrescis;\n",
		"vita detestabilis\nnunc obdurat\net tunc curat\n",
		"ludo mentis aciem,\negestatem,\npotestatem\ndissolvit ut glaciem.\n",
	}
	ft, err := ofortuna.NewFortuneTeller(strings.NewReader(oFortuna))
	if err != nil {
		t.Fatal(err)
	}
	var got []string = ft.Fortune
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestNewFortuneTellerFromFile(t *testing.T) {
	t.Parallel()
	var ft ofortuna.FortuneTeller
	ft, err := ofortuna.NewFortuneTellerFromFile("testdata/quotes.fortune")
	if err != nil {
		t.Fatal(err)
	}
	wantLen := 10
	gotLen := len(ft.Fortune)
	if wantLen != gotLen {
		t.Fatalf("want %d fortunes, got %d", wantLen, gotLen)
	}
	wantFortune := "What is a master but a master student?\n—Neil Peart"
	gotFortune := ft.Fortune[2]
	if wantFortune != gotFortune {
		t.Fatalf("want %q, got %q", wantFortune, gotFortune)
	}
}