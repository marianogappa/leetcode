/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {
        if (head == NULL) {
            return false;
        }
        ListNode *slow = head;
        ListNode *fast = head;
        if (head->next == NULL) {
            return false;
        }
        fast = head->next;
        while (slow != NULL && fast != NULL && slow != fast) {
            slow = slow->next;
            fast = fast->next;
            if (fast != NULL) {
                fast = fast->next;
            }
        }
        if (slow == NULL || fast == NULL) {
            return false;
        }
        return true;
    }
};

