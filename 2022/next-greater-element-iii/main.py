from typing import Optional

# It's trivial but has a million tricky parts.
#
# Time: O(1)
# Space: O(1)
#
# Intuition: from right to left flip the first decreasing digit with the first digit larger than it (again r-to-l).
#            Then, sort ascending all numbers after that first descending digit.
#
# Example: 2147483476
#          2147483[4]76   first decreasing digit is 4
#          2147483[4]7[6] first digit larger than it is 6
#          2147483[6]7[4] flip them
#          21474836[74]   all numbers after first descending digit
#          21474836[47]   sort them ascending
#          2147483647     done! Remember to make sure it doesn't exceed 2^31-1
class Solution:
    def nextGreaterElement(self, n: int) -> int:
        # Getting the reversed digits makes everything easier
        reversed_n = get_reversed_digits(n)

        # Flip the first decreasing digit with the first digit larger than it
        idx = flip_first_decreasing_digit(reversed_n)
        if idx is None:
            return -1

        # Sort ascending all numbers up to first decreasing digit (here descending because it's reversed)
        reversed_n = sorted(reversed_n[:idx], reverse=True) + reversed_n[idx:]

        # We now have the number we want (there's no smaller one). But make sure it doesn't exceed 2^31-1
        # without accidentally using math with numbers larger than 2^31-1
        reversed_max = get_reversed_digits(2**31 - 1)
        if len(reversed_n) == len(reversed_max) and reverse(reversed_n) > reverse(
            reversed_max
        ):
            return -1

        # Number is ready. Just need to go from reversed digits to int.
        return reversed_digits_to_int(reversed_n)


def get_reversed_digits(n: int) -> list[int]:
    digits = []
    while n > 0:
        digits.append(n % 10)
        n //= 10

    return digits


def reverse(ns: list[int]) -> list[int]:
    return ns[::-1]


def reversed_digits_to_int(ns: list[int]) -> int:
    num = 0
    multiplier = 1
    for n in ns:
        num += multiplier * n
        multiplier *= 10

    return num


def flip_first_decreasing_digit(ns: list[int]) -> Optional[int]:
    for i in range(1, len(ns)):
        # Found first decreasing digit at i
        if ns[i] < ns[i - 1]:
            # Now, find smallest digit larger than ns[i]
            for j in range(0, i):
                if ns[j] > ns[i]:
                    # Swap them
                    ns[i], ns[j] = ns[j], ns[i]
                    return i

    return None


print(Solution().nextGreaterElement(12))
print(Solution().nextGreaterElement(21))
print(Solution().nextGreaterElement(13))
print(Solution().nextGreaterElement(333333))
print(Solution().nextGreaterElement(333331))
print(Solution().nextGreaterElement(333334))
print(Solution().nextGreaterElement(2147483647))
print(Solution().nextGreaterElement(230241))
print(Solution().nextGreaterElement(2147483476))
