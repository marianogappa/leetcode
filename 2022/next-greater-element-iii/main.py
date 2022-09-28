from typing import Optional


class Solution:
    def nextGreaterElement(self, n: int) -> int:
        reversed_n = get_reversed_digits(n)
        reversed_max = get_reversed_digits(2**31 - 1)

        idx = flip_first_decreasing_digit(reversed_n)
        if idx is None:
            return -1

        reversed_n = sorted(reversed_n[:idx], reverse=True) + reversed_n[idx:]

        if len(reversed_n) == len(reversed_max) and reverse(reversed_n) > reverse(
            reversed_max
        ):
            return -1

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


# 2147483476
