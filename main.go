package main

import "fmt"

func main() {
	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	trx := 2
	fmt.Println(maxProfit(trx, prices))
	fmt.Println("==========================================")

	trx = 3
	fmt.Println(maxProfit(trx, prices))
	fmt.Println("==========================================")

	trx = 4
	fmt.Println(maxProfit(trx, prices))
}

func maxProfit(trx int, prices []int) int {
	minCosts := make([]int, len(prices))
	minCosts[0] = prices[0]
	maxProfits := make([]int, len(prices))
	currTrx := 1
	res := 0

	if len(prices) == 0 {
		return 0
	}

	for i := 0; i < trx; i++ {
		for j := 1; j < len(prices); j++ {
			tempCost := prices[j] - maxProfits[j]
			minCosts[j] = min(minCosts[j-1], tempCost)

			tempProfit := prices[j] - minCosts[j]
			maxProfits[j] = max(maxProfits[j-1], tempProfit)
		}

		currProfit := maxProfits[len(maxProfits)-1]
		if res == currProfit {
			break
		} else {
			res = currProfit
			currTrx++
		}
	}

	fmt.Println("Max Transaki yang didapat berada pada Transaksi ke", currTrx-1, ". Dengan total:", res)

	return res
}
