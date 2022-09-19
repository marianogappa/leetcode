class Solution:
    # Time: O(s)
    # Space: O(s)
    #
    # This algoritm is incredibly simple but it's a brainf*** to figure out why it works.
    #
    # First, intuition on how to solve the exercise:
    #
    # The shortest palindrome is the one that adds the minimum prefix to s, right?
    #
    # e.g. "radar" is a palindrome. For "radarhi", adding a "ih": "ih[radar]hi".
    # e.g. "hiradar" => "radari[h]iradar".
    #
    # The strategy is to find the longest palindrome that starts at s[0], and prepend the reverse of the rest of s.
    # Unfortunately, answering "longest palindrome that starts at s[0]" is quadratic.
    #
    # Remember the KMP table? It calculates the longest prefix of the pattern that equals the suffix up at the ith
    # character, and it's calculated in linear time.
    #
    # So consider "radarhi|ihradar": the last number on the KMP table would answer the longest prefix that equals the
    # suffix at the end of the string, that is, that "matches its reverse": so the longest palindrome starting at 0!
    #
    # Note that if a character is not added in the middle, the result could potentially be larger than "s".
    #
    # Once we know the longest prefix, just prepend the part of the reverse that is not part of the palindrome, to "s".
    def shortestPalindrome(self, s: str) -> str:
        reverse_s = s[::-1]
        lps = calculate_longest_prefix_suffix(f"{s}|{reverse_s}")
        len_of_longest_palindrome_from_start = lps[-1]
        reverse_of_non_palindrome_part = reverse_s[
            :-len_of_longest_palindrome_from_start
        ]
        return reverse_of_non_palindrome_part + s


# This is just the verbatim KMP table algorithm, no modifications
def calculate_longest_prefix_suffix(pattern: str) -> list[int]:
    lps = [0] * len(pattern)
    for i in range(1, len(pattern)):
        ps_len = lps[i - 1]

        while ps_len > 0 and pattern[i] != pattern[ps_len]:
            ps_len = lps[ps_len - 1]

        if pattern[i] == pattern[ps_len]:
            ps_len += 1

        lps[i] = ps_len

    return lps


print(Solution().shortestPalindrome("aacecaaa"), "== aaacecaaa")
print(Solution().shortestPalindrome("abcd"), "== dcbabcd")
