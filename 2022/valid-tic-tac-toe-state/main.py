# wip!
# It's straightforward, but just make sure to spend sufficient time
# looking at the possible rules and look for counterexamples.
from collections import Counter
# Time: O(1)
# Space: O(1)
class Solution:
    def validTicTacToe(self, board: List[str]) -> bool:
        freqs = Counter(f"{board[0]}{board[1]}{board[2]}")
        
        o_wins = (
            int(board[0][0] == board[0][1] == board[0][2] == "O") +
            int(board[1][0] == board[1][1] == board[1][2] == "O") +
            int(board[2][0] == board[2][1] == board[2][2] == "O") +

            int(board[0][0] == board[1][0] == board[2][0] == "O") +
            int(board[0][1] == board[1][1] == board[2][1] == "O") +
            int(board[0][2] == board[1][2] == board[2][2] == "O") +

            int(board[0][0] == board[1][1] == board[2][2] == "O") +
            int(board[2][0] == board[1][1] == board[0][2] == "O")
        )
            
        x_wins = (
            int(board[0][0] == board[0][1] == board[0][2] == "X") +
            int(board[1][0] == board[1][1] == board[1][2] == "X") +
            int(board[2][0] == board[2][1] == board[2][2] == "X") +

            int(board[0][0] == board[1][0] == board[2][0] == "X") +
            int(board[0][1] == board[1][1] == board[2][1] == "X") +
            int(board[0][2] == board[1][2] == board[2][2] == "X") +

            int(board[0][0] == board[1][1] == board[2][2] == "X") +
            int(board[2][0] == board[1][1] == board[0][2] == "X")
        )

        # Both cannot win
        if o_wins and x_wins:
            return False
        
        # There cannot be more Os than Xs
        if freqs['O'] > freqs['X']:
            return False
        
        # If O wins, there must be equal number of Os and Xs
        if o_wins and freqs['O'] != freqs['X']:
            return False
        
        # If X wins, there must be one more X than O
        if x_wins and freqs['O'] + 1 != freqs['X']:
            return False

        # There can be 0 or 1 Xs more than Os
        if freqs['X'] > freqs['O'] + 1:
            return False
        
        return True
