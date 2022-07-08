package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dog struct {
	weight int
}

func (d dog) feedPerMonth() int {
	const feedPerKg int = 2
	return feedPerKg * d.weight
}

func (d dog) Weight() int {
	return d.weight
}

func (d dog) String() string {
	return "dog"
}

type cat struct {
	weight int
}

func (c cat) feedPerMonth() int {
	const feedPerKg int = 7
	return feedPerKg * c.weight
}

func (c cat) Weight() int {
	return c.weight
}

func (c cat) String() string {
	return "cat"
}

type cow struct {
	weight int
}

func (c cow) feedPerMonth() int {
	const feedPerKg int = 25
	return feedPerKg * c.weight
}

func (c cow) Weight() int {
	return c.weight
}

func (c cow) String() string {
	return "cow"
}

type animals interface {
	feedPerMonthGetter
	WeightGetter
	fmt.Stringer
}

type feedPerMonthGetter interface {
	feedPerMonth() int
}

type WeightGetter interface {
	Weight() int
}

func animalInfo(a []animals) (totalFeedForFarm int) {
	for _, v := range a {
		fmt.Printf("This is a %s, it weights %v kg, and it needs %v kg of feed per month.\n", v.String(), v.Weight(), v.feedPerMonth())
		totalFeedForFarm += v.feedPerMonth()
	}
	return totalFeedForFarm
}

func generateAnimals(n int) []animals {
	var farmAnimals []animals

	const minCatWeight int = 1
	const maxCatWeight int = 20

	const minDogWeight int = 1
	const maxDogWeight int = 90

	const minCowWeight int = 50
	const maxCowWeight int = 600

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		newRandomAnimal := rand.Intn(3)
		switch newRandomAnimal {
		case 0:
			farmAnimals = append(farmAnimals, dog{weight: rand.Intn(maxDogWeight) + minDogWeight})
		case 1:
			farmAnimals = append(farmAnimals, cat{weight: rand.Intn(maxCatWeight) + minCatWeight})
		case 2:
			farmAnimals = append(farmAnimals, cow{weight: rand.Intn(maxCowWeight) + minCowWeight})
		}
	}
	return farmAnimals
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randIntOfAnimals := rand.Intn(20) + 3

	farmAnimals := generateAnimals(randIntOfAnimals)

	total := animalInfo(farmAnimals)
	fmt.Printf("Total feed per mounth %v kg.\n", total)
}
