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
	BookProduction float64
	TotalBooks     float64
}

var (
	rate          float64
	weeks         int
	reveUptime    int
	fokUptime     int
	uniUptime     int
	useReve       bool
	useUni        bool
	useFoK        bool
	uniPercentage float64
)

var scientists []Scientist
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	setup()

	var totalScientists = 5
	var ticks = weeks * 7 * 24
	var sciGen float64 = 0

	var reveRandom, uniRandom int

	for i := 0; i <= ticks; i++ {
		reveRandom = r.Intn(100) + 1
		uniRandom = r.Intn(100) + 1

		switch {
		case !useReve && !useUni:
			sciGen += rate
		case useReve && !useUni:
			if reveRandom >= reveUptime {
				sciGen += rate
			} else {
				sciGen += (rate * 1.3)
			}
		case !useReve && useUni:
			if uniRandom > uniUptime {
				sciGen += (rate)

			} else { // And Universities are active this tick
				sciGen += (rate * (1 + (calcUniScientistSpawnRate(uniPercentage, 0) / 100)))
			}

		case useReve && useUni:
			// If Revelation is not active this tick...
			if reveRandom >= reveUptime {
				// And Universities are inactive this tick
				if uniRandom > uniUptime {
					sciGen += (rate)

				} else { // And Universities are active this tick
					sciGen += (rate * (1 + (calcUniScientistSpawnRate(uniPercentage, 0) / 100)))
				}
			} else { // Revelation is active this tick...
				// And Universities are not active
				if uniRandom > uniUptime {
					sciGen += (rate * 1.3)
				} else { // Universities are active
					sciGen += (rate * 1.3 * (1 + (calcUniScientistSpawnRate(uniPercentage, 0) / 100)))
				}
			}
		}

		// If science gen = 100 or more
		if sciGen >= 100 {
			sciGen = sciGen - 100
			totalScientists += 1
			addScientist("basic", 0)
		}
		incrementScientistLifetime()
		incrementScientistTotalBooks(uniRandom)
		checkScientistRank()
	}

	fmt.Println("Total Scientists: " + strconv.Itoa(len(scientists)))
	fmt.Printf("Total Books Generated: %.0f\n", calcTotalBooks())

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

func incrementScientistTotalBooks(uniRandom int) {
	var fokRandom = r.Intn(100)

	for i := 0; i < len(scientists); i++ {
		switch {
		// Both FoK and Universities are not used this tick
		case !useFoK && !useUni:
			scientists[i].TotalBooks += (scientists[i].BookProduction)
		// FoK is possible but Universities are not active
		case useFoK && !useUni:
			if fokRandom > fokUptime {
				scientists[i].TotalBooks += (scientists[i].BookProduction)
			} else {
				scientists[i].TotalBooks += (scientists[i].BookProduction * 1.1)
			}
		case !useFoK && useUni:
			if uniRandom > uniUptime {
				scientists[i].TotalBooks += scientists[i].BookProduction
			} else { // and we do have universities active this tick
				scientists[i].TotalBooks += (scientists[i].BookProduction * (1 + (calcUniBookProduction(uniPercentage, 0) / 100)))
			}
		case useFoK && useUni:
			// If FoK is not active...
			if fokRandom > fokUptime {
				// And we do not have universities this tick
				if uniRandom > uniUptime {
					scientists[i].TotalBooks += scientists[i].BookProduction
				} else { // and we do have universities active this tick
					scientists[i].TotalBooks += (scientists[i].BookProduction * (1 + (calcUniBookProduction(uniPercentage, 0) / 100)))
				}
			} else { // FoK is always active
				// And we do not have univeristies this tick
				if uniRandom > uniUptime {
					fmt.Println("here")

					scientists[i].TotalBooks += scientists[i].BookProduction * 1.1
				} else {
					fmt.Println("here2")
					scientists[i].TotalBooks += (scientists[i].BookProduction * (1 + (calcUniBookProduction(uniPercentage, 0) / 100)) * 1.1)
				}
			}
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

func calcTotalBooks() float64 {
	var totalBooks float64 = 0

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

	if err != nil {
		fmt.Println(err.Error())
	}

	weeks = config.WeeksInAge
	rate = config.ScientistGenerationRate
	reveUptime = config.RevelationDownTime
	fokUptime = config.FoundtainOfKnowledgeDowntime
	uniUptime = config.UniversityUptime
	uniPercentage = config.UniversityPercentage

	if reveUptime != 0 {
		useReve = true
	}

	if uniUptime != 0 {
		useUni = true
	}

	if fokUptime != 0 {
		useFoK = true
	}

	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)
	addScientist("basic", 0)

}

var config *configStruct

type configStruct struct {
	ScientistGenerationRate         float64 `json : "ScientistGenerationRate"`
	WeeksInAge                      int     `json:"WeeksInAge"`
	RevelationDownTime              int     `json:"RevelationUptime"`
	FoundtainOfKnowledgeDowntime    int     `json:"FountainOfKnowledgeUpime"`
	UniversityPercentage            float64 `json:"UniversityPercentage"`
	RacialScienceProductionModifier float64 `json:"RacialScienceProductionModifier"`
	UniversityUptime                int     `json:"UniversityUptime"`
}
