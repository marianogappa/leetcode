class Solution:
    # Time: O(n)
    # Space: O(n)
    #
    # The intuitive solution is to sort the citations reverse, but it's n*logn time.
    #
    # A more clever solution is that since len(citations) <= 5000, we can use linear space
    # and bucket sort papers by citations up to len(citations).
    #
    # At that point, just traverse the buckets in reverse order until max-h is found.
    def hIndex(self, citations: List[int]) -> int:
        buckets = [0] * (len(citations)+1)

        # For each paper, bucket it by citations (up to len(citations))
        for count in citations:
            buckets[min(count, len(citations))] += 1

        # Go from larger to lower bucket, and sum papers with h+ citations
        paper_sum = 0
        for h in range(len(buckets)-1, -1, -1):
            paper_sum += buckets[h]

            # As soon as the sum of papers with h+ citations is >= h, max h has been found.
            if paper_sum >= h:
                return h

        # Unreachable
