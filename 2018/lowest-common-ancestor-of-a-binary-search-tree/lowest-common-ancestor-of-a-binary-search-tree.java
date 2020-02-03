/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        ArrayList<TreeNode> l = dfs(root, p, new ArrayList<>());
        ArrayList<TreeNode> r = dfs(root, q, new ArrayList<>());
        for (int i = 0; i < l.size(); i++) {
            if (i == r.size() || l.get(i) != r.get(i)) {
                return l.get(i - 1);
            }
        }
        return l.get(l.size() - 1);
    }

    private ArrayList<TreeNode>dfs(TreeNode c, TreeNode e, ArrayList<TreeNode> ns) {
        if (c == null) {
            return new ArrayList<>();
        }
        ns.add(c);
        if (c == e) {
            return ns;
        }
        ArrayList<TreeNode> lns = dfs(c.left, e, new ArrayList<>(ns));
        if (lns.size() > 0) {
            return lns;
        }
        ArrayList<TreeNode> rns = dfs(c.right, e, new ArrayList<>(ns));
        if (rns.size() > 0) {
            return rns;
        }
        return new ArrayList<>();
    }
}
