# Code is pretty self-explanatory.
# Trickiest bit is that to find the number to swap to, you have
# to search from the right to produce the highest number.
#
# Think like this: to swap to highest number, take the small number
# from the left, and the high number from the right. This will
# find the largest number.

# Time: O(1)
# Space: O(1)
class Solution:
    def maximumSwap(self, num: int) -> int:
        digits = number_to_digits(num)
        sorted_digits = sorted(digits, reverse=True)

        idx = first_differing_index(digits, sorted_digits)

        if idx == -1:
            return num

        right_idx = find_from_right(digits, sorted_digits[idx])

        digits = swap(digits, idx, right_idx)

        return digits_to_int(digits)


def swap(l: list[int], i: int, j: int) -> list[int]:
    l[i], l[j] = l[j], l[i]
    return l


def first_differing_index(l1: list[int], l2: list[int]) -> int:
    for i in range(len(l1)):
        if l1[i] != l2[i]:
            return i
    return -1


def find_from_right(digits: list, target_num: int) -> int:
    j = len(digits)-1
    while j >= 0:
        if target_num == digits[j]:
            return j
        j -= 1
    return 0


def digits_to_int(digits: list[int]):
    number = 0
    multiplier = 1
    j = len(digits)-1
    while j >= 0:
        number += digits[j] * multiplier
        multiplier *= 10
        j -= 1
    return number


def number_to_digits(num: int) -> list[int]:
    digits = []
    while num > 0:
        digits.append(num % 10)
        num //= 10
    digits = digits[::-1]
    return digits
