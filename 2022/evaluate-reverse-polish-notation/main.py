# Mostly trivial: store operands in a stack and consume them upon an operator.
#
# Tricky parts:
# - Division is NOT `//`! It's `int(a / b)`!!
# - Note that popping 2 operands pops them in reverse order to how they are used.
# - Remember to PUSH the result of each operation as a new operand in the stack.

class Solution:
    # Time: O(n)
    # Space: O(n)
    def evalRPN(self, tokens: List[str]) -> int:
        funcs = {
            '+': lambda a, b: a + b,
            '-': lambda a, b: a - b,
            '*': lambda a, b: a * b,
            '/': lambda a, b: int(a / b),
        }
        operands = []
        for token in tokens:
            if token in ['+', '-', '*', '/']:
                op_2 = operands.pop()
                op_1 = operands.pop()
                operands.append(funcs[token](op_1, op_2))
            else:
                operands.append(int(token))
        
        return operands.pop()
