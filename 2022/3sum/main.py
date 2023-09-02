# Producing all triplets is `n^3`, but we can produce all pairs in `n^2` and complete the
# triplet in constant time by having a dict of idxs of every number.
#
# This solution TLEs, so optimise by keeping only 3- idxs of each number. Any more would
# produce a dupe triplet.
#
# To avoid dupe triplets, we can sort the triplet and add it as a tuple to a set.

# Time: O(n^2) where n is the len of nums since we produce all pairs
# Space: O(n) since we store all nums in dict, list, set & output
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # First keep idxs of every number, to complete triplets in constant time
        num_to_idxs = defaultdict(list)
        deduped_nums = []
        for i, num in enumerate(nums):
            # Optimisation: keep up to 3 idxs of each number (only dupe triplets after)
            if len(num_to_idxs[num]) < 3:
                deduped_nums.append(num)
                num_to_idxs[num].append(len(deduped_nums)-1)
        
        triplets = set([])
        # Then produce every pair of numbers (n[i], n[j]) (quadratic step!)
        for i in range(len(deduped_nums)):
            for j in range(i + 1, len(deduped_nums)):
                # Complete the triplet with an n[k] in constant time
                for k in num_to_idxs.get(-(deduped_nums[i] + deduped_nums[j]), []):
                    # Use only i < j < k to ensure unique idxs
                    if k > j:
                        # Sort triplet and add it as a tuple on a set so we can de-dupe
                        triplet = sorted([deduped_nums[i], deduped_nums[j], deduped_nums[k]])
                        triplets.add((triplet[0], triplet[1], triplet[2]))
        
        # Output format must be list of lists
        return [list(triplet) for triplet in triplets]
