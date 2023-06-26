package best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		result int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, d := range cases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, d.result, maxProfit(d.prices))
		})
	}
}
