# wip!
# Very simple if you figure out the trick.
#
# Pick a random word and get the guesses. If not done, the secret word must have EQUAL matches to this random guess.
#
# Use this knowledge to build a new word list with only the words that do. Then repeat. You'll find it.

# """
# This is Master's API interface.
# You should not implement it, or speculate about its implementation
# """
# class Master:
#     def guess(self, word: str) -> int:
from random import randint

# Time: O(a*n) where n is the whole content of words, and a is the number of allowed guesses (as an upper bound).
# Space: O(a*n) same logic, as we create a new word list every time.
class Solution:
    def findSecretWord(self, words: List[str], master: 'Master') -> None:
        # This is just to not loop forever
        while len(words):
            # Make a random guess, and see how many matches it got against secret word
            guess = randint(0, len(words)-1)
            guess_matches = master.guess(words[guess])

            # Build word list of words that are not the guess and have EQUAL matches with guess
            next_words = []
            for i in range(len(words)):
                if i != guess and matches(words[guess], words[i]) == guess_matches:
                    next_words.append(words[i])

            words = next_words

# Count equal characters between two words
def matches(w1: str, w2: str) -> int:
    return sum([1 if w1[i] == w2[i] else 0 for i in range(6)])
