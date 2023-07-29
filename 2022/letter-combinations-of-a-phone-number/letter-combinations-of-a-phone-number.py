# Think of each digit as a node, and each mapped letter as a child. Then, we're just doing a BFS on a tree!
#
# To calculate Time complexity, we just need to know how many nodes are in the tree.
#
# There are d levels (one per digit), and each node has 3 or 4 children, so that's 4^d nodes.
#
# Time: O(d*4^d) where d is the number of digits (times the linear append operation)
# Space: O(4^d) because on BFS we store one level at a time, and last level has half of the nodes
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []

        digit_to_letters: dict[str, list[str]] = {
            "2": ["a", "b", "c"],
            "3": ["d", "e", "f"],
            "4": ["g", "h", "i"],
            "5": ["j", "k", "l"],
            "6": ["m", "n", "o"],
            "7": ["p", "q", "r", "s"],
            "8": ["t", "u", "v"],
            "9": ["w", "x", "y", "z"],
        }

        combinations = digit_to_letters[digits[0]]
        for i in range(1, len(digits)): # level of tree
            new_combinations = []
            for letter in digit_to_letters[digits[i]]:
                for i in range(len(combinations)):
                    new_combinations.append(combinations[i] + letter)  # this append is linear to string size

            combinations = new_combinations

        return combinations
