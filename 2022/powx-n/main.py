# This makes perfect sense once you think in binary!
#
# For 2 ^ 10:
# 1) make sure n is positive. If not, just do 1/x and n = -1.
# 2) Convert the exponent to binary: for 10 it's "1010"
# 3) Note that 2^10 = 2^8 * 2^4. The exponents coincide with the "1s" in the binary representation.

# Time: O(logn)
# Space: O(1)
class Solution:
    def myPow(self, x, n):
        # Normalize n to positive
        if n < 0:
            x = 1 / x
            n = -n

        result = 1
        while n:
            # Check if the rightmost bit in n is 1
            if n % 2 == 1:
                result *= x
            
            # Every time we advance a bit is like squaring the number
            x *= x

            # Shift right, so that we can check the next bit to the right
            n >>= 1

        return result
