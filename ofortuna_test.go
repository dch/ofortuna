package ofortuna_test

import (
	"testing"

	"ofortuna"
)

func TestGetRandomFortunes(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		want string
	}{
		{name: "returns a string", want: ofortuna.O_Fortuna},
	}

	for _, tc := range testCases {
		got := ofortuna.GetRandomFortune()
		if tc.want != got {
			t.Errorf("%v: wanted %v, got %v", tc.name, tc.want, got)
		}
	}
}
