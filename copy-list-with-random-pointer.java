import java.util.ArrayList;

class RandomListNode {
    int label;
    RandomListNode next, random;
    RandomListNode(int x) { this.label = x; }
};

public class Solution {
    public RandomListNode copyRandomList(RandomListNode head) {
        if (head == null) {
            return null;
        }
        Map<RandomListNode, Integer> m = new HashMap<RandomListNode, Integer>();
        RandomListNode cur = head;
        int i = 0;
        List<RandomListNode> cp = new ArrayList<RandomListNode>();
        while (cur != null) {
            m.put(cur, i);
            cp.add(new RandomListNode(cur.label));
            i++;
            cur = cur.next;
        }
        cur = head;
        i = 0;
        while (cur != null) {
            if (i > 0) {
                cp.get(i - 1).next = cp.get(i);
            }
            if (m.containsKey(cur.random)) {
                cp.get(i).random = cp.get(m.get(cur.random));
            }
            cur = cur.next;
            i++;
        }
        return cp.get(0);
    }
}
