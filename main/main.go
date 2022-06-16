package main

import "fmt"

const (
	applePrice float64 = 5.99
	pearPrice  float64 = 7
	ourMoney   float64 = 23
)

func main() {
	applesToBuy := 9
	pearsToBuy := 8
	price := (applePrice * float64(applesToBuy)) + (pearPrice * float64(pearsToBuy))
	fmt.Printf("1. We need %.2f UAH to buy %v apples and %v pears.\n", price, applesToBuy, pearsToBuy)

	weCanBuyPears := ourMoney / pearPrice
	fmt.Printf("2. We can buy %v pears.\n", int(weCanBuyPears))

	weCanBuyApples := float64(ourMoney) / applePrice
	fmt.Printf("3. We can buy %v apples.\n", int(weCanBuyApples))

	wantApples := 2
	wantPears := 2
	priceForApples := applePrice * float64(wantApples)
	priceForPears := pearPrice * float64(wantPears)
	canBuy := priceForApples+priceForPears <= float64(ourMoney)
	fmt.Printf("4. Can we buy %v pears and %v apples? %t \n", wantApples, wantPears, canBuy)
}
