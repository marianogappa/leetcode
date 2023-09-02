# wip!
class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = parse_path(path)
        return "/" + "/".join(stack)

def parse_path(path: str) -> list[str]:
    stack = []
    for part in path.split('/'):
        if part == "" or part == ".":
            continue
        elif part == ".." and len(stack) > 0:
            stack.pop()
        elif part != "..":
            stack.append(part)
    
    return stack

print(Solution().simplifyPath("/home/"), "== /home")
print(Solution().simplifyPath("/../"), "== /")
print(Solution().simplifyPath("/home//foo/"), "== /home/foo")
