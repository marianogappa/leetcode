# wip!
class Solution:
    def minCostSetTime(self, startAt: int, moveCost: int, pushCost: int, targetSeconds: int) -> int:
        opts = options(targetSeconds)
        min_cost = float("inf")
        for opt in opts:
            min_cost = min(min_cost, calculate_cost(
                startAt, moveCost, pushCost, get_digits(opt)))
        return min_cost


def calculate_cost(startAt: int, moveCost: int, pushCost: int, digits: list[int]) -> int:
    if not digits:
        return 0
    cur_move_cost = moveCost if startAt != digits[0] else 0
    return cur_move_cost + pushCost + calculate_cost(digits[0], moveCost, pushCost, digits[1:])


def options(targetSeconds: int) -> list[int]:
    divisor = targetSeconds // 60
    options = []
    for i in range(divisor, -1, -1):
        rest = targetSeconds - 60*i
        if 100*i+rest > 9999:
            continue
        if rest > 99:
            break
        options.append(100*i+rest)
    return options


def get_digits(num: int) -> list[int]:
    digits: list[int] = []
    while num > 0:
        digits.append(num % 10)
        num //= 10
    return reverse(digits)


def reverse(digits: list[int]) -> list[int]:
    i = 0
    j = len(digits)-1
    while i < j:
        digits[i], digits[j] = digits[j], digits[i]
        i += 1
        j -= 1
    return digits


print(Solution().minCostSetTime(
    0,
    100000,
    100000,
    6039,
))
