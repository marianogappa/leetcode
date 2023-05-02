# Exercise is straightforward but has a million edge cases: key is naming things properly.
#
# Analyze every line:
# - On files, path length should be checked and a max len kept.
# - On directories, a length of the current directory structure should be calculated and appended into a stack.
# - On both, how many dirs it went back, forward must be calculated, and the stack must be updated. It is incredibly
#   tricky to figure out how to do this properly.

# Time: O(n)
# Space: O(n)
class Solution:
    def lengthLongestPath(self, input: str) -> int:
        max_len: int = 0
        dir_lens: list[int] = [0]

        for entry in input.split("\n"):
            tab_count, is_file, entry_len = analyze_entry(entry)

            remove_tab_count = len(dir_lens) - tab_count
            if remove_tab_count > 0 and len(dir_lens):
                dir_lens = dir_lens[:-remove_tab_count]

            cur_dir_len = dir_lens[-1] if len(dir_lens) else 0

            if is_file:
                max_len = max(max_len, cur_dir_len + entry_len)
            else:
                dir_lens.append(cur_dir_len + entry_len)

        return max_len


def analyze_entry(s: str) -> tuple[int, bool, int]:
    tab_count = count_character("\t", s)
    is_file = count_character(".", s) > 0
    entry_len = len(s) - tab_count if is_file else len(s) - tab_count + 1
    return tab_count, is_file, entry_len


def count_character(needle, haystack: str) -> int:
    count = 0
    for char in haystack:
        if char == needle:
            count += 1
    return count


print(Solution().lengthLongestPath("dir\nother_dir\nfile.e"), "== 6")
print(Solution().lengthLongestPath("dir\nother_dire\n\tfile.e"), "== 17")
print(Solution().lengthLongestPath("file.path"), "== 9")
print(Solution().lengthLongestPath("dir"), "== 0")
print(Solution().lengthLongestPath("dir\n\tfile.e"), "== 10")
print(Solution().lengthLongestPath("dir\n\tfile.e\nfile_with_longer_size.a"), "== 23")
print(Solution().lengthLongestPath("dir\n\tfile.e\nfi.a"), "== 10")
print(Solution().lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"), "== 20")
print(
    Solution().lengthLongestPath(
        "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
    ),
    "== 32",
)
print(Solution().lengthLongestPath("a\n\tb1\n\t\tf1.txt\n\taaaaa\n\t\tf2.txt"), "== 14")
