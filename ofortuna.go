package ofortuna

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

type fortuneSet struct {
	fortune []string
}

func GetRandomFortune() string {
	return O_Fortuna
}
