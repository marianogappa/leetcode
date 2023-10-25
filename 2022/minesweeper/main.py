DELTAS = [
    [0,1],
    [1,0],
    [-1,0],
    [0,-1],
    [1,1],
    [-1,-1],
    [1,-1],
    [-1,1],
]

class Solution:
    def updateBoard(self, board: List[List[str]], click: List[int]) -> List[List[str]]:
        x, y = click[1], click[0]

        if board[y][x] == "M":
            board[y][x] = "X"
            return board
        
        dfs(board, x, y)
        return board

def dfs(board, x, y):
    if y < 0 or x < 0 or y >= len(board) or x >= len(board[0]):
        return
    
    if board[y][x] != 'E':
        return

    adj_mine_count = count_adjacent_mines(board, x, y)
    if adj_mine_count:
        board[y][x] = str(adj_mine_count)
        return
    
    board[y][x] = 'B'
    for delta in DELTAS:
        dx, dy = delta
        px, py = x+dx, y+dy
        dfs(board, px, py)

def count_adjacent_mines(board, x, y) -> int:
    count = 0
    for delta in DELTAS:
        dx, dy = delta
        px, py = x+dx, y+dy
        if py < 0 or px < 0 or py >= len(board) or px >= len(board[0]):
            continue    
        if board[py][px] in ['M', 'X']:
            count += 1

    return count
