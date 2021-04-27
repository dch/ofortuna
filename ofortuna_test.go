package ofortuna_test

import (
	"testing"

	"ofortuna"
)

func TestGetRandomFortunes(t *testing.T) {
	t.Parallel()
	var got string = ofortuna.GetRandomFortune()
	if len(got) == 0 {
		t.Errorf("want non-empty fortune")
	}
}

func TestNewFortuneTellerReadsFortunes(t *testing.T) {
	// TODO construct input reader from testdata/quotes.fortune
	ft, err := ofortuna.NewFortuneTeller(input)
	if err != nil {
		t.Fatal(err)
	}
	wantLen := 10
	gotLen := ft.Len()
	if wantLen != gotLen {
		t.Errorf("want %d fortunes, got %d", wantLen, gotLen)
	}

}
