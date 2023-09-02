# wip!
# Straightforward if using proper abstractions. For a parseable grammar, not too complex to implement a token iterator.
# Only 3 tokens exist: an int, a string or a closing bracket (opening bracket always comes after number so ignore it!)
#
# Switch:
#
# If string: append it to the solution as is.
# If number: get the number, run this decoding function recursively to get the string, and append n*decoded to solution.
# If closing bracket: we must be in a recursive call! End the call here.
# If no tokens left: we must be in a non-recursive call! End the call here.

from dataclasses import dataclass
from typing import Optional

# Time: O(n) we parse the string once. There are constants in this string that will make us for-loop over it.
# Space: O(n) we store the solution in a buffer (in this case actually in an array); we also store partial strings.
class Solution:
    def decodeString(self, s: str) -> str:
        return do_decode(TokenIterator(s))

@dataclass
class Token:
    token_type: str
    int_value: int = 0
    str_value: str = ""


class TokenIterator:
    def __init__(self, s: str):
        self.s = s
        self.i = 0
    
    def next(self) -> Optional[Token]:
        if self.i >= len(self.s):
            return None

        if self.s[self.i] == "]":
            self.i += 1
            return Token(token_type="]")

        if self.s[self.i] >= "0" and self.s[self.i] <= "9":
            token = []
            while self.i < len(self.s) and self.s[self.i] >= "0" and self.s[self.i] <= "9":
                token.append(self.s[self.i])
                self.i +=1
            self.i += 1 # Ignore always present opening bracket
            return Token(token_type="int", int_value=int(''.join(token)))

        if self.s[self.i] >= "a" and self.s[self.i] <= "z":
            token = []
            while self.i < len(self.s) and self.s[self.i] >= "a" and self.s[self.i] <= "z":
                token.append(self.s[self.i])
                self.i +=1
            return Token(token_type="str", str_value=''.join(token))
    
def do_decode(it: TokenIterator) -> str:
    result: list[str] = []

    while (token := it.next()) and token.token_type != "]":

        if token.token_type == "str":
            result.append(token.str_value)

        if token.token_type == "int":
            result.append(do_decode(it)*token.int_value)
    
    return ''.join(result)

print(Solution().decodeString('3[a]2[bc]') == 'aaabcbc')
print(Solution().decodeString('3[a2[c]]') == 'accaccacc')
print(Solution().decodeString('2[abc]3[cd]ef') == 'abcabccdcdcdef')
