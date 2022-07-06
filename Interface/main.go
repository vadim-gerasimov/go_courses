package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dog struct {
	feedPerKg int
	weight    int
}

func (d dog) feedPerMonth() int {
	return d.feedPerKg * d.weight
}

func (d dog) animalWeight() int {
	return d.weight
}

func (d dog) String() string {
	return "dog"
}

type cat struct {
	feedPerKg int
	weight    int
}

func (c cat) feedPerMonth() int {
	return c.feedPerKg * c.weight
}

func (c cat) animalWeight() int {
	return c.weight
}

func (c cat) String() string {
	return "cat"
}

type cow struct {
	feedPerKg int
	weight    int
}

func (c cow) feedPerMonth() int {
	return c.feedPerKg * c.weight
}

func (c cow) animalWeight() int {
	return c.weight
}

func (c cow) String() string {
	return "cow"
}

type animals interface {
	feedPerMonthGetter
	animalWeightGetter
	fmt.Stringer
}

type feedPerMonthGetter interface {
	feedPerMonth() int
}

type animalWeightGetter interface {
	animalWeight() int
}

func animalInfo(a []animals) int {
	totalFeedForFarm := 0
	for _, v := range a {
		fmt.Printf("This is a %s, it weights %v, and it needs %v of feed per month.\n", v.String(), v.animalWeight(), v.feedPerMonth())
		totalFeedForFarm += v.feedPerMonth()
	}
	return totalFeedForFarm
}

func generateAnimals(n int) []animals {
	var farmAnimals []animals

	minCatWeight := 1
	maxCatWeight := 20

	minDogWeight := 1
	maxDogWeight := 90

	minCowWeight := 50
	maxCowWeight := 600

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		newRandomAnimal := rand.Intn(3)
		switch newRandomAnimal {
		case 0:
			farmAnimals = append(farmAnimals, dog{weight: int(rand.Intn(maxDogWeight) + minDogWeight), feedPerKg: 2})
		case 1:
			farmAnimals = append(farmAnimals, cat{weight: int(rand.Intn(maxCatWeight) + minCatWeight), feedPerKg: 7})
		case 2:
			farmAnimals = append(farmAnimals, cow{weight: int(rand.Intn(maxCowWeight) + minCowWeight), feedPerKg: 25})
		}
	}
	return farmAnimals
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randIntOfAnimals := rand.Intn(20) + 3

	farmAnimals := generateAnimals(randIntOfAnimals)

	total := animalInfo(farmAnimals)
	fmt.Printf("Total feed per mounth %v\n", total)
}
