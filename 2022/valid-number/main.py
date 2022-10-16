class Solution:
    """
        Time: O(n)
        Space: O(n)

        If the "in" with set is not optimal, a Trie can be used to make the time linear.
        This solution relies on converting the string to a list of tokens, which makes the space linear.
        With the Trie solution, each time a token is popped, one could traverse the Trie rather than store the token,
        making the space constant.


        Valid Number formats:
            
        1) Integer
        
        +1
        1
        
        2) Decimal Number:
            
        +1.
        1.
        +1.1
        1.1
        +.1
        .1

        3) Integer, 'e' or 'E', Integer:
        
        +1e1
        1e1

        +1e+1
        1e+1
        
        4) Decimal Number, 'e' or 'E', Integer:
            
        +1.e1
        1.e1
        +1.1e1
        1.1e1
        +.1e1
        .1e1
        
        +1.e+1
        1.e+1
        +1.1e+1
        1.1e+1
        +.1e+1
        .1e+1
    """
    def isNumber(self, s: str) -> bool:
        parser = Parser(s)
        tokens = []
        while (token := parser.pop()) != "end":
            if token == "error":
                return False
            tokens.append(token)
        
        valid_number_formats = {
            "+1",
            "1",
            "+1.",
            "1.",
            "+1.1",
            "1.1",
            "+.1",
            ".1",
            "+1e1",
            "1e1",
            "+1e+1",
            "1e+1",
            "+1.e1",
            "1.e1",
            "+1.1e1",
            "1.1e1",
            "+.1e1",
            ".1e1",
            "+1.e+1",
            "1.e+1",
            "+1.1e+1",
            "1.1e+1",
            "+.1e+1",
            ".1e+1",
        }

        return ''.join(tokens) in valid_number_formats
        

class Parser:
    def __init__(self, s: str):
        self.s = s
        self.i = 0
    
    def pop(self) -> tuple[str]:
        if self.i >= len(self.s):
            return "end"
        if self.s[self.i] >= "0" and self.s[self.i] <= "9":
            while self.i < len(self.s) and self.s[self.i] >= "0" and self.s[self.i] <= "9":
                self.i += 1
            return "1"
        if self.s[self.i] == "+" or self.s[self.i] == "-":
            self.i += 1
            return "+"
        if self.s[self.i] == "e" or self.s[self.i] == "E":
            self.i += 1
            return "e"
        if self.s[self.i] == ".":
            self.i += 1
            return "."
        return "error"

positives = [Solution().isNumber(example) for example in ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]]
print(positives)

negatives = [Solution().isNumber(example) for example in ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]]
print(negatives)
