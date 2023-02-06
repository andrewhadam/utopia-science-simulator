package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var (
	rate         float64
	weeks        int
	reveDowntime int
	fokDowntime  int
)

var scientists []Scientist
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Represents % of tiem that Fountain of Knowledge will not be up

func main() {

	setup()

	var numOfScientists = 5
	var ticks = weeks * 7 * 24
	var sciGen float64 = 0
	var adjustedRate = rate * 1

	var ranNum int

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
	n := Scientist{Rank: "Recruit", Lifetime: 0, BookProduction: 70, TotalBooks: 0}
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
		if scientists[i].TotalBooks > 1679 && scientists[i].TotalBooks < 5520 {
			scientists[i].BookProduction = 80
			scientists[i].Rank = "Novice"
		}
		if scientists[i].TotalBooks > 5519 && scientists[i].TotalBooks < 12000 {
			scientists[i].BookProduction = 90
			scientists[i].Rank = "Graduate"
		}
		if scientists[i].TotalBooks > 11999 {
			scientists[i].BookProduction = 100
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

	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())

	}

	err = json.Unmarshal(file, &config)
	fmt.Println(file)

	if err != nil {
		fmt.Println(err.Error())
	}

	weeks = config.WeeksInAge
	rate = config.ScientistGenerationRate
	reveDowntime = config.RevelationDownTime
	fokDowntime = config.FoundtainOfKnowledgeDowntime

	fmt.Println(weeks)
	fmt.Println(rate)
	fmt.Println(reveDowntime)
	fmt.Println(fokDowntime)

	//	Token = config.Token
	//	BotPrefix = config.BotPrefix

	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)

}

var config *configStruct

type configStruct struct {
	ScientistGenerationRate      float64 `json : "ScientistGenerationRate"`
	WeeksInAge                   int     `json:"WeeksInAge"`
	RevelationDownTime           int     `json:"RevelationDowntime"`
	FoundtainOfKnowledgeDowntime int     `json:"FountainOfKnowledgeDowntime"`
}
