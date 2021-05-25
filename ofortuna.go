package ofortuna

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

type fortuneTeller struct {
	Fortune []string
	Random *rand.Rand
}

func NewFortuneTeller(r io.Reader) (fortuneTeller, error) {
	var ft = fortuneTeller{
		Fortune: []string{},
	}
	scanner := bufio.NewScanner(r)
	var fortune, line string
	for scanner.Scan() {
		line = scanner.Text()
		if line != "%" {
			fortune += line + "\n"
			continue
		}
		ft.Fortune = append(ft.Fortune, fortune)
		fmt.Printf("fortune %q\n", fortune)
		fortune = ""
	}
	if scanner.Err() != nil {
		return fortuneTeller{}, scanner.Err()
	}
	ft.Fortune = append(ft.Fortune, fortune)
	return ft, nil
}

func (fortuneTeller) GetRandomFortune() string {
	return ""
}