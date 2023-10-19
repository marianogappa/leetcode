# Merging accounts means grouping the sets of emails if they have an email in common.
#
# This is textbook UnionFind, where vertex => email & edges => grouped emails in each account.
#
# There's a little trickery to producing the output format (check comments on solution).

class Solution:
    # Time: O(n) note that UF has path compression so ops are ~O(1), sort is O(1) too
    # Space: O(n)
    def accountsMerge(self, accounts: List[List[str]]) -> List[List[str]]:
        # Build a mapping from all emails to their name.
        email_to_name = {}
        for account in accounts:
            for email in account[1:]:
                email_to_name[email] = account[0]

        # Add all emails to the UF data structure.
        uf = UnionFind()
        for account in accounts:
            for email in account[1:]:
                uf.add(email)
        
        # Union all emails within each account.
        for account in accounts:
            for i in range(2, len(account)):
                uf.union(account[i-1], account[i])
        
        # N.B. At this point, we have solved the exercise, but still need to build result DS.

        # Build a mapping from group to all emails.
        # Note that the group key will be a meaningless but UNIQUE id from UF DS.
        groups = defaultdict(list)
        for email in email_to_name.keys():
            groups[uf.find(email)].append(email)
        
        # Result is a list[list[str]]. Emails must be sorted.
        # First entry on each inner list must be a name: use the initial mapping for this.
        result = []
        for _, emails in groups.items():
            result.append([email_to_name[emails[0]], *sorted(emails)])
        
        return result


class UnionFind:
    parents: list[int]
    idxs: dict[any, int]
    sizes: list[int]

    def __init__(self):
        self.parents = []
        self.idxs = {}
        self.sizes = []
    
    def add(self, vertex: any) -> int:
        idx = len(self.parents)
        self.parents.append(idx)
        self.sizes.append(1)
        self.idxs[vertex] = idx
    
    def union(self, v1: any, v2: any) -> None:
        p1 = self.find(v1)
        p2 = self.find(v2)
        if self.sizes[p1] > self.sizes[p2]:
            self.sizes[p1] += self.sizes[p2]
            self.parents[p2] = p1
        else:
            self.sizes[p2] += self.sizes[p1]
            self.parents[p1] = p2
    
    def _find(self, idx: int) -> int:
        return idx if self.parents[idx] == idx else self._find(self.parents[idx])
    
    def find(self, vertex: any) -> int:
        return self._find(self.idxs[vertex])
