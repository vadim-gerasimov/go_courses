package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type animals interface {
	feedPerMonthGetter
	weightGetter
	nameGetter
	edibilityGetter
	typeGetter
}

type nameGetter interface {
	nameOf() string
}

type typeGetter interface {
	typeOf() string
}
type feedPerMonthGetter interface {
	feedPerMonth() int
}

type weightGetter interface {
	weightOf() int
}

type edibilityGetter interface {
	edibilityOf() bool
}

func animalInfo(a []animals) (totalFeedForFarm int, err error) {
	for _, v := range a {
		err = validation(v)
		if err != nil {
			switch {
			case errors.Is(err, typeErr):
				fmt.Printf("for animal type of %s and name %s: %s", v.typeOf(), v.nameOf(), err)
				continue
			case errors.Is(err, normalWeightErr):
				fmt.Printf("for animal type of %s and name %s: %s", v.typeOf(), v.nameOf(), err)
				continue
			case errors.Is(err, edibilityErr):
				err = fmt.Errorf("for animal type of %s and name %s: %w", v.typeOf(), v.nameOf(), err)
				return 0, err
			default:
				err = errors.New("undeclareted err")
				return 0, err
			}
		}
		fmt.Printf("This is a %s, it weights %v kg, and it needs %v kg of feed per month.\n", v.nameOf(), v.weightOf(), v.feedPerMonth())
		totalFeedForFarm += v.feedPerMonth()
	}
	return totalFeedForFarm, nil
}

func validation(a animals) error {
	if err := validateTypeAssertion(a); err != nil {
		err = fmt.Errorf("failed type validation: %w", err)
		return err
	}
	if err := validateNormalWeight(a); err != nil {
		err = fmt.Errorf("failed normal weight validation: %w", err)
		return err
	}
	if err := validateEdibility(a); err != nil {
		err = fmt.Errorf("failed edibility validation: %w", err)
		return err
	}
	return nil
}

var unknownBeastErr = errors.New("unknown beast")
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
	default:
		return unknownBeastErr
	}
	if a.nameOf() != animalType {
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
	default:
		return unknownBeastErr
	}
	if a.weightOf() < normalWeight {
		return fmt.Errorf("%w: must weight more than %d kg\n", normalWeightErr, normalWeight)
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
	default:
		return unknownBeastErr
	}
	if a.edibilityOf() != animalEdibility {
		return fmt.Errorf("%w: %s edibility must be %t but is not\n", edibilityErr, a.typeOf(), animalEdibility)
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
		return
	}
	fmt.Printf("\nTotal feed per mounth %v kg.\n", total)

}
