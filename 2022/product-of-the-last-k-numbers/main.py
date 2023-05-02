# By inspecting an example it looks obvious that it's a division:
#
# Consider [1, 2, 3, 4]:
#
# Calculate running product:
# 1, 1*2, 1*2*3, 1*2*3*4
#
# If at that point you're asked for k = 2, then answer is 3*4,
# which can be calculated as accum[3]/accum[3-k]:
#
# 1*2*3*4/(1*2) = 3*4
#
# It's almost as easy as that. The only tricky part is that
# zeros break the product, so they cannot be included in the
# running product.
#
# There's a simple solution for that though: just start the
# running product again after a zero, and deal with negative
# index overflows as an edge case.

# Space: O(n)
class ProductOfNumbers:

    def __init__(self):
        self.accum: list[int] = []

    # Time: O(1)
    def add(self, num: int) -> None:
        if num == 0:
            self.accum = []
        elif len(self.accum) == 0:
            self.accum.append(num)
        else:
            self.accum.append(num*self.accum[len(self.accum)-1])

    # Time: O(1)
    def getProduct(self, k: int) -> int:
        last_i = len(self.accum)-1
        divide_i = last_i - k
        if divide_i == -1:
            return self.accum[last_i]
        elif divide_i <= -2:
            return 0
        else:
            return self.accum[last_i]//self.accum[divide_i]
