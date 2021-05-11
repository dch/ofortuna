package ofortuna_test

import (
	"math/rand"
	"ofortuna"
	"strings"
	"testing"
)

func TestGetRandomFortunes(t *testing.T) {
	t.Parallel()
	var fortunes = strings.NewReader("foo bar baz")
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
