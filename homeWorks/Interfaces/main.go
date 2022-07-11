package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type dog struct {
	weight    int
	name      string
	edibility bool
}

func (d dog) feedPerMonth() int {
	const feedPerKg int = 2
	return feedPerKg * d.weight
}

func (d dog) Weight() int {
	return d.weight
}

func (d dog) String() string {
	return d.name
}

func (d dog) Edibility() bool {
	return d.edibility
}

func (d dog) Type() string {
	return "dog"
}

type cat struct {
	weight    int
	name      string
	edibility bool
}

func (c cat) feedPerMonth() int {
	const feedPerKg int = 7
	return feedPerKg * c.weight
}

func (c cat) Weight() int {
	return c.weight
}

func (c cat) String() string {
	return c.name
}

func (c cat) Edibility() bool {
	return c.edibility
}
func (c cat) Type() string {
	return "cat"
}

type cow struct {
	weight    int
	name      string
	edibility bool
}

func (c cow) feedPerMonth() int {
	const feedPerKg int = 25
	return feedPerKg * c.weight
}

func (c cow) Weight() int {
	return c.weight
}

func (c cow) String() string {
	return c.name
}

func (c cow) Edibility() bool {
	return c.edibility
}

func (c cow) Type() string {
	return "cow"
}

type animals interface {
	feedPerMonthGetter
	WeightGetter
	NameGetter
	edibilityGetter
	TypeGetter
}

type NameGetter interface {
	String() string
}

type TypeGetter interface {
	Type() string
}
type feedPerMonthGetter interface {
	feedPerMonth() int
}

type WeightGetter interface {
	Weight() int
}

type edibilityGetter interface {
	Edibility() bool
}

func animalInfo(a []animals) (totalFeedForFarm int, err error) {
	for _, v := range a {
		err = validation(v)
		if err != nil {
			if errors.Is(err, typeErr) {
				fmt.Printf("for animal type of %s and name %s: %s", v.Type(), v.String(), err)
				continue
			}
			if errors.Is(err, normalWeightErr) {
				fmt.Printf("for animal type of %s and name %s: %s", v.Type(), v.String(), err)
				continue
			}
			if errors.Is(err, edibilityErr) {
				err = fmt.Errorf("for animal type of %s and name %s: %w", v.Type(), v.String(), err)
				return 0, err
			}
		}
		fmt.Printf("This is a %s, it weights %v kg, and it needs %v kg of feed per month.\n", v.String(), v.Weight(), v.feedPerMonth())
		totalFeedForFarm += v.feedPerMonth()
	}
	return totalFeedForFarm, nil
}

func validation(a animals) error {
	err := validateTypeAssertion(a)
	if err != nil {
		err = fmt.Errorf("failed type validation: %w", err)
		return err
	}
	err = validateNormalWeight(a)
	if err != nil {
		err = fmt.Errorf("failed normal weight validation: %w", err)
		return err
	}
	err = validateEdibility(a)
	if err != nil {
		err = fmt.Errorf("failed edibility validation: %w", err)
		return err
	}
	return nil
}

var typeErr = errors.New("wrong type")

func validateTypeAssertion(a animals) error {
	var animalType string
	switch a.(type) {
	case dog:
		animalType = "dog"
	case cat:
		animalType = "cat"
	case cow:
		animalType = "cow"
	}
	if a.String() != animalType {
		return fmt.Errorf("%w: must be %s but is not\n", typeErr, animalType)
	}
	return nil
}

var normalWeightErr = errors.New("underweight")

func validateNormalWeight(a animals) error {
	var normalWeight int
	switch a.(type) {
	case dog:
		normalWeight = minDogWeight
	case cat:
		normalWeight = minCatWeight
	case cow:
		normalWeight = minCowWeight
	}
	if a.Weight() < normalWeight {
		return fmt.Errorf("%w: must weight more than %d\n", normalWeightErr, normalWeight)
	}
	return nil
}

var edibilityErr = errors.New("wrong edibility")

func validateEdibility(a animals) error {
	var animalEdibility bool
	switch a.(type) {
	case dog:
		animalEdibility = false
	case cat:
		animalEdibility = false
	case cow:
		animalEdibility = true
	}
	if a.Edibility() != animalEdibility {
		return fmt.Errorf("%w: %s edibility must be %t but is not\n", edibilityErr, a.Type(), animalEdibility)
	}
	return nil
}

const (
	minDogWeight int = 5
	minCatWeight int = 5
	minCowWeight int = 50
	maxDogWeight int = 90
	maxCatWeight int = 20
	maxCowWeight int = 600
)

func generateAnimals(n int) []animals {
	var farmAnimals []animals

	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		newRandomAnimal := rand.Intn(3)
		switch newRandomAnimal {
		case 0:
			farmAnimals = append(farmAnimals, dog{weight: rand.Intn(maxDogWeight), name: "dog", edibility: false})
		case 1:
			farmAnimals = append(farmAnimals, cat{weight: rand.Intn(maxCatWeight), name: "cat", edibility: false})
		case 2:
			farmAnimals = append(farmAnimals, cow{weight: rand.Intn(maxCowWeight), name: "cow", edibility: true})
		}
	}
	farmAnimals = append(farmAnimals, cow{weight: rand.Intn(maxCowWeight), name: "cw", edibility: true})
	farmAnimals = append(farmAnimals, cow{weight: 60, name: "cow", edibility: false})

	return farmAnimals
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randIntOfAnimals := rand.Intn(20) + 3

	farmAnimals := generateAnimals(randIntOfAnimals)

	total, err := animalInfo(farmAnimals)
	if err != nil {
		fmt.Printf("critical err: %v", err)
	} else {
		fmt.Printf("\nTotal feed per mounth %v kg.\n", total)
	}
}
