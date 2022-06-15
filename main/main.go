package main

import "fmt"

const (
	applePrice float64 = 5.99
	pearPrice  int     = 7
	ourMoney   int     = 23
)

func main() {
	applesToBuy := 9
	pearsToBuy := 8
	price := (applePrice * float64(applesToBuy)) + float64(pearPrice*pearsToBuy)
	fmt.Printf("1. We need %.2f UAH to buy %v apples and %v pears.\n", price, applesToBuy, pearsToBuy)

	weCanBuyPears := ourMoney / pearPrice
	fmt.Printf("2. We can buy %v pears.\n", weCanBuyPears)

	weCanBuyApples := float64(ourMoney) / applePrice
	fmt.Printf("3. We can buy %v apples.\n", int(weCanBuyApples))

	wantApples := 2
	wantPears := 2
	canBuy := float64(pearPrice*wantPears)+(applePrice*float64(wantApples)) <= float64(ourMoney)
	fmt.Printf("4. Can we buy %v pears and %v apples? %t \n", wantApples, wantPears, canBuy)
}
