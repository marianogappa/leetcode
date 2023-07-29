# Resist the intuitive solution to try DP based on the coins: "I'll try to use the larger coins first to reach the target with less coins".
#
# Consider the case of coins = [1, 3, 4] and amount = 6: this algo would use 4 + 1 + 1, but optimal is to use 3 + 3.
#
# Instead, use DP on the amounts, bottom up. For each amount, try to use each coin, and greedily keep the optimal solution.
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        # Initialise memo with a large number except for the base case of trying to reach 0 amount, which takes 0 coins
        minCoins = [0] * (amount+1)
        for i in range(len(minCoins)):
            minCoins[i] = float("inf")
        minCoins[0] = 0

        # For every amount ranging 1..amount:
        for amt in range(1, amount+1):
            # For every coin denomination:
            for coin in coins:
                # If subtracting current coin we surpass the target amount, can't use it.
                # If subtracting current coin we arrive a target amount with no solution, can't use it.
                # (Note that, by now, we must know optimal solutions for smaller amounts!)
                if amt - coin < 0 or minCoins[amt-coin] == float("inf"):
                    continue

                # We encountered a way to reach the current amount with a number of coins, but it may
                # not yet be the optimal solution, so greedily check all possible ones and keep the 
                # one with the smallest number of coins.
                minCoins[amt] = min(minCoins[amt], 1+minCoins[amt-coin])

        # If by the end of this exercise there isn't an option for the specified amount, 
        # then there isn't one.
        if minCoins[amount] == float("inf"):
            return -1

        # Otherwise return it.
        return minCoins[amount]
