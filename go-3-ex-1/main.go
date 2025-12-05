package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

type FullName struct {
	FirstName string
	LastName  string
}
type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'

	if p.Month == 1 {
		if p.Day <= 19 {
			zodiacSign = Capricornus
		} else {
			zodiacSign = Aquarius
		}
	} else if p.Month == 2 {
		if p.Day <= 18 {
			zodiacSign = Aquarius
		} else {
			zodiacSign = Pisces
		}
	} else if p.Month == 3 {
		if p.Day <= 20 {
			zodiacSign = Pisces
		} else {
			zodiacSign = Aries
		}
	} else if p.Month == 4 {
		if p.Day <= 19 {
			zodiacSign = Aries
		} else {
			zodiacSign = Taurus
		}
	} else if p.Month == 5 {
		if p.Day <= 20 {
			zodiacSign = Taurus
		} else {
			zodiacSign = Gemini
		}
	} else if p.Month == 6 {
		if p.Day <= 20 {
			zodiacSign = Gemini
		} else {
			zodiacSign = Cancer
		}
	} else if p.Month == 7 {
		if p.Day <= 22 {
			zodiacSign = Cancer
		} else {
			zodiacSign = Leo
		}
	} else if p.Month == 8 {
		if p.Day <= 22 {
			zodiacSign = Leo
		} else {
			zodiacSign = Virgo
		}
	} else if p.Month == 9 {
		if p.Day <= 22 {
			zodiacSign = Virgo
		} else {
			zodiacSign = Libra
		}
	} else if p.Month == 10 {
		if p.Day <= 22 {
			zodiacSign = Libra
		} else {
			zodiacSign = Scorpius
		}
	} else if p.Month == 11 {
		if p.Day <= 21 {
			zodiacSign = Scorpius
		} else {
			zodiacSign = Sagittarius
		}
	} else if p.Month == 12 {
		if p.Day <= 21 {
			zodiacSign = Sagittarius
		} else {
			zodiacSign = Capricornus
		}
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	dennis := Person{FullName{"Dennis", "Ritchie"}, BirthDate{9, 9, 1941}}
	rick := Person{FullName{"Rick", "Astley"}, BirthDate{6, 2, 1966}}
	edsger := Person{FullName{"Edsger", "Dijkstra"}, BirthDate{11, 5, 1930}}
	alan := Person{FullName{"Alan", "Turing"}, BirthDate{23, 6, 1912}}

	outputWithZodiacSign(grace)
	outputWithZodiacSign(dennis)
	outputWithZodiacSign(rick)
	outputWithZodiacSign(edsger)
	outputWithZodiacSign(alan)
}
