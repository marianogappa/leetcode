class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        i = 0

        stack = []
        while i < len(s):
            # Pop
            char = s[i]
            start = i
            while i < len(s) and s[i] == char and i - start + 1 <= k:
                i += 1
            count = i - start

            if count == k:
                continue
            if not stack or stack[-1][0] != char:
                stack.append((char, count))
            else:
                _, old_count = stack.pop()
                new_count = (old_count + count) % k
                if new_count > 0:
                    stack.append((char, new_count))
        
        return ''.join([char*count for char, count in stack])

