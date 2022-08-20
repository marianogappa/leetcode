from collections import defaultdict


class Solution:
    # Time: O(n+m) technically worst case intersection n*m
    # Space: O(n+m) storing dict sets
    #
    # Intuition is to build freq maps of both, but because
    # indices matter, instead we build a dict from number to
    # a set of indices. The intersection of sets provides the
    # bulls, and the cows will be the non-bulls in those sets,
    # but there can be as many as the smallest len of the sets,
    # minus bulls.
    #
    # In principle it seems linear time & space, but remember
    # that the worst case intersection of sets could be n*m
    # with very bad hashset structure, if every "val in big_set"
    # is linear instead of constant.
    def getHint(self, secret: str, guess: str) -> str:
        secret_dict_set = build_dict_set(secret)
        guess_dict_set = build_dict_set(guess)

        bulls = cows = 0
        for num, guess_set in guess_dict_set.items():
            secret_set = secret_dict_set[num]
            num_bulls = len(secret_set.intersection(guess_set))
            cows += min(len(secret_set), len(guess_set)) - num_bulls
            bulls += num_bulls

        return f"{bulls}A{cows}B"


def build_dict_set(nums: str) -> dict[set[str]]:
    dict_set = defaultdict(set[int])
    for i, num in enumerate(nums):
        dict_set[num].add(i)
    return dict_set
