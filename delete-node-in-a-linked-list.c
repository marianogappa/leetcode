#include <stdio.h>

struct ListNode {
  int val;
  struct ListNode *next;
};

void deleteNode(struct ListNode* node) {
  node->val = node->next->val;
  node->next = node->next->next;
}

int main() {

}
