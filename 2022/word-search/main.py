# If the starting point was fixed (e.g. from top-left), finding the word would be
# running DFS on the decision tree of looking for the next letter in all 4 directions
# recursively.
#
# Since it's a graph (not a tree), we must keep a visited set and crucially backtrack
# if a path doesn't work.
#
# Since the starting point is NOT fixed, simply run this exact algorithm from every cell!

class Solution:
    # Time: O(x * y * 4^w) because we'll DFS from scratch from every cell
    # Space: O(x * y + w) to store the set and the stack of max depth `w`
    def exist(self, board: List[List[str]], word: str) -> bool:
        # Try from every starting cell, using DFS
        return any(
            exists_from(x, y, 0, set([]), board, word) 
            for y in range(len(board)) 
            for x in range(len(board[0]))
        )

# Time: O(4^w), because we'll DFS on a tree of depth len(word), with 4 children each
def exists_from(x: int, y: int, i: int, visited: set[tuple[int, int]], board: List[List[str]], word: str) -> bool:
    # If we reached the end of the word, we found it!
    if i == len(word):
        return True

    # If we're out of bounds and haven't reached the end of the word, this path didn't work
    if x < 0 or y < 0 or y >= len(board) or x >= len(board[0]):
        return False

    # If the current letter doesn't match expected, this path doesn't work
    if board[y][x] != word[i]:
        return False

    # If we already visited this cell, this path doesn't work
    if (x, y) in visited:
        return False
    
    # This path is promising, so let's visit this cell, and try next...
    visited.add((x, y))

    # From the current cell, visit all neighbouring for the next character in word
    for delta in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
        if exists_from(x+delta[0], y+delta[1], i+1, visited, board, word):
            return True
    
    # IMPORTANT! This path didn't work, so let's backtrack!
    visited.remove((x, y))
    return False
