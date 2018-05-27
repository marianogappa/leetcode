class Solution:
	def maxProfit(self, prices):
		"""
		:type prices: List[int]
		:rtype: int
		"""
		buy_day = 0
		max_profit = 0
		for i in range(len(prices)):
			if prices[i] < prices[buy_day]:
				buy_day = i
			max_profit = max(max_profit, prices[i] - prices[buy_day])
		return max_profit
