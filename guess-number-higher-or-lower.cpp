// Forward declaration of guess API.
// @param num, your guess
// @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
int guess(int num);

class Solution {
public:
    int guessNumber(int n) {
        int upper = n;
        int lower = 0;
        int g = (upper - lower) / 2 + lower;
        int r = guess(g);
        while (r != 0) {
            if (r == 1) {
                lower = g + 1;
            } else if (r == -1) {
                upper = g - 1;
            }
            g = (upper - lower) / 2 + lower;
            r = guess(g);
        }
        return g;
    }
};
