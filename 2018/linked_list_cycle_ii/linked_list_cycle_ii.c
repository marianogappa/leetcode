#include <stdio.h>

struct ListNode {
  int val;
  struct ListNode *next;
};

struct ListNode *detectCycle(struct ListNode *head) {
  struct ListNode *slow = head;
  struct ListNode *fast = head;
  struct ListNode *entry = head;

  while (slow != NULL && fast != NULL && fast->next != NULL && fast->next->next != NULL) {
    slow = slow->next;
    fast = fast->next->next;

    if (slow == fast) {
      while (slow != entry) {
        slow = slow->next;
        entry = entry->next;
      }
      return entry;
    }
  }
  return NULL;
}

int main() {
  struct ListNode *test1, *test2, *test3, *test4, *test5, *test6, *test7;
  test6 = &(struct ListNode) {6, NULL};
  test5 = &(struct ListNode) {5, test6};
  test4 = &(struct ListNode) {4, test5};
  test3 = &(struct ListNode) {3, test4};
  test2 = &(struct ListNode) {2, test3};
  test1 = &(struct ListNode) {1, test2};
  test6->next = test5;

  printf("(%d)\n", detectCycle(test1)->val);
}
