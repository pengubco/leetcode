package best_time_to_buy_and_sell_stock_ii

/*
Key observation: Because there is no limitations on number of transactions, so there is no difference between
(i,j)+(j,k) and (i,k). So the greedy algorithm works.
1. Whenever a[i] > a[i-1], make a trade (i-1, i).
2. Whenever a[i] <= a[i-1], ignore a[i-1], because buying a[i] is always better at buying at a[i-1].
*/

func maxProfit(prices []int) int {
	result := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}
