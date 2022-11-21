package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Scientist struct {
	Rank           string
	Lifetime       int
	BookProduction float32
	TotalBooks     float32
}

var rate float64 = 1.5
var weeks int = 10
var scientists []Scientist
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Represents % of time that Revelation will not be up
var reveDowntime = 40

// Represents % of tiem that Fountain of Knowledge will not be up
var fokDowntime = 40

func main() {

	var numOfScientists = 5
	var ticks = weeks * 7 * 24
	var sciGen float64 = 0
	var adjustedRate = rate * 1

	var ranNum int

	setup()

	for i := 0; i <= ticks; i++ {
		ranNum = r.Intn(100)
		fmt.Printf("Revelation Chance: %d\n", ranNum)

		if ranNum >= 0 && ranNum < reveDowntime {
			sciGen += adjustedRate
			fmt.Printf("SciGen False: %f\n", sciGen)
		} else {
			sciGen += (adjustedRate * 1.3)
			fmt.Printf("SciGen True: %f\n", sciGen)
		}

		// sciGen += adjustedRate

		// If science gen = 100 or more
		if sciGen >= 100 {
			sciGen = sciGen - 100
			numOfScientists += 1
			addScientist("basic", 0)
		}
		incrementScientistLifetime()
		incrementScientistTotalBooks()
		checkScientistRank()
	}

	fmt.Println(math.Mod(float64(ticks)*rate, 100))
	fmt.Println((float64(ticks) * rate / 100))
	fmt.Println(numOfScientists)

	fmt.Println("Total Scientists: " + strconv.Itoa(len(scientists)))
	fmt.Printf("Total Books Generated: %f\n", calcTotalBooks())
	fmt.Printf("University Scientist Spawn Rate: %.2f\n", calcUniScientistSpawnRate(18, 0))

}

func addScientist(r string, l int) {
	n := Scientist{Rank: "Recruit", Lifetime: 0, BookProduction: 100, TotalBooks: 0}
	scientists = append(scientists, n)
}

func incrementScientistLifetime() {
	for i := 0; i < len(scientists); i++ {
		scientists[i].Lifetime += 1
	}
}

func incrementScientistTotalBooks() {
	for i := 0; i < len(scientists); i++ {
		var ranNum = r.Intn(100)
		if ranNum >= 0 && ranNum < fokDowntime {
			scientists[i].TotalBooks += scientists[i].BookProduction
		} else {
			scientists[i].TotalBooks += (scientists[i].BookProduction * 1.1)
		}
	}
}

func checkScientistRank() {
	for i := 0; i < len(scientists); i++ {
		if scientists[i].TotalBooks > 2399 && scientists[i].TotalBooks < 8160 {
			scientists[i].BookProduction = 120
			scientists[i].Rank = "Novice"
		}
		if scientists[i].TotalBooks > 8159 && scientists[i].TotalBooks < 18240 {
			scientists[i].BookProduction = 140
			scientists[i].Rank = "Graduate"
		}
		if scientists[i].TotalBooks > 18239 {
			scientists[i].BookProduction = 160
			scientists[i].Rank = "Professor"
		}
	}
}

func calcTotalBooks() float32 {
	var totalBooks float32 = 0

	for i := 0; i < len(scientists); i++ {
		totalBooks += scientists[i].TotalBooks
	}
	return totalBooks
}

func calcUniScientistSpawnRate(percent float64, racialMod float64) float64 {
	return (2 * math.Min(50, (percent*(1+racialMod))) * (100 - (math.Min(50, (percent * (1 + racialMod)))))) / 100
}

func calcUniBookProduction(percent float64, racialMod float64) float64 {
	return (1 * math.Min(50, (percent*(1+racialMod))) * (100 - (math.Min(50, (percent * (1 + racialMod)))))) / 100
}

func setup() {
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
}
