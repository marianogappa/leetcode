# Easy but too many edge cases. Just figure out the wording algorithm. There are special
# cases for <=10, <=20, and the functions for <100 and <1000 can be reused by larger number
# implementations. After 1000, there's a pattern: [2] Million [345] Thousand [123].
#
# Special cases: 0 is "Zero", but 100 is not "[One] Hundred [Zero]".
# A number group that is zero should not be translated e.g. 134000123 => [134] Million [1] Hundred [23]

# Time: O(1)
# Space: O(1)
class Solution:
    def numberToWords(self, num: int) -> str:
        if num < 1000:
            return ' '.join(until1000(num, no_zero=False))
        
        threes = [
            [],
            ["Thousand"],
            ["Million"],
            ["Billion"],
            ["Trillion"],
            ["Quadrillion"],
            ["Quintillion"],
            ["Sixtillion"],
            ["Septillion"],
            ["Octillion"],
            ["Nonillion"],
        ]

        snum = str(num)

        # Group by threes right to left
        groups = [int(snum[max(i-2, 0):i+1]) for i in range(len(snum)-1, -1, -3)]

        parts = []
        for i, group in enumerate(groups):
            if group == 0:
                continue
            parts += [' '.join(until1000(group, no_zero=True) + threes[i])]
        
        return ' '.join(reversed(parts))

def until1000(num: int, *, no_zero: bool) -> list[str]:
    if num < 100:
        return until100(num, no_zero=no_zero)
    
    hundreds, rest = int(str(num)[0]), int(str(num)[1:])

    return until100(hundreds, no_zero=True) + ["Hundred"] + until100(rest, no_zero=True)

def until100(num: int, *, no_zero: bool) -> list[str]:
    if num == 0 and no_zero:
        return []

    digits = [
        "Zero",
        "One",
        "Two",
        "Three",
        "Four",
        "Five",
        "Six",
        "Seven",
        "Eight",
        "Nine",
    ]
    
    if num < 10:
        return [digits[num]]
    
    special = [
        "Ten",
        "Eleven",
        "Twelve",
        "Thirteen",
        "Fourteen",
        "Fifteen",
        "Sixteen",
        "Seventeen",
        "Eighteen",
        "Nineteen",
    ]

    if num < 20:
        return [special[num-10]]
    
    tens = [
        "Zero",
        "Ten",
        "Twenty",
        "Thirty",
        "Forty",
        "Fifty",
        "Sixty",
        "Seventy",
        "Eighty",
        "Ninety",
    ]

    first_digit = num % 10
    num //= 10
    second_digit = num % 10

    return [tens[second_digit]] + until100(first_digit, no_zero=True)


print(Solution().numberToWords(10))
print(Solution().numberToWords(100))
print(Solution().numberToWords(1))
print(Solution().numberToWords(0))
print(Solution().numberToWords(13))
print(Solution().numberToWords(20))
print(Solution().numberToWords(203))
print(Solution().numberToWords(253))
print(Solution().numberToWords(2053))
print(Solution().numberToWords(1234567))
