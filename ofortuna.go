package ofortuna

import (
	"bufio"
	"io"
	"math/rand"
	"time"
)

type FortuneTeller struct {
	Fortune []string
	Random *rand.Rand
}

func NewFortuneTeller(r io.Reader) (FortuneTeller, error) {
	var ft = FortuneTeller{
		Fortune: []string{},
		Random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var fortune, line string
	for scanner.Scan() {
		line = scanner.Text()
		if line != "%" {
			fortune += line + "\n"
			continue
		}
		ft.Fortune = append(ft.Fortune, fortune)
		fortune = ""
	}
	if scanner.Err() != nil {
		return FortuneTeller{}, scanner.Err()
	}
	ft.Fortune = append(ft.Fortune, fortune)
	return ft, nil
}

func (ft FortuneTeller) GetRandomFortune() string {
	index := ft.Random.Intn(len(ft.Fortune))
	return ft.Fortune[index]
}