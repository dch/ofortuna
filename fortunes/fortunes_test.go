package fortunes_test

import (
	"testing"

	"github.com/dch/ofortuna/fortunes"
)

func TestGetRandomFortunes(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		want string
	}{
		{name: "returns a string", want: fortunes.O_Fortuna},
	}

	for _, tc := range testCases {
		got := fortunes.GetRandomFortune()
		if tc.want != got {
			t.Errorf("%v: wanted %v, got %v", tc.name, tc.want, got)
		}
	}
}
