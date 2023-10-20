# The most straightforward solution seems to simply go step by step.

class Solution:
    def myAtoi(self, s: str) -> int:
        # Ignore any leading whitespace
        i = 0
        while i < len(s) and s[i] == ' ':
            i += 1
        
        # Return zero if there are no more chars
        if i == len(s):
            return 0
        
        # Sign is optional. Consume it if it exists.
        sign = 1
        if s[i] == '+':
            i += 1
        elif s[i] == '-':
            sign = -1
            i += 1
        
        # Return zero if there are no more chars
        if i == len(s):
            return 0

        # Read all digits
        str_to_digit = {
            '0': 0,
            '1': 1,
            '2': 2,
            '3': 3,
            '4': 4,
            '5': 5,
            '6': 6,
            '7': 7,
            '8': 8,
            '9': 9,
        }
        digits = []
        while i < len(s) and str_to_digit.get(s[i]) is not None:
            digits.append(str_to_digit[s[i]])
            i += 1
        
        # Pop each digit and push it as each digit in num
        mul = 1
        num = 0
        while digits:
            digit = digits.pop()
            num += digit * mul
            mul *= 10
        
        # Apply the clamp logic for positives and negatives
        if sign == 1 and num > 2**31 -1:
            return 2**31 -1
        
        if sign == -1 and num > 2**31:
            return -2**31

        # Don't forget to apply the sign
        return num * sign
