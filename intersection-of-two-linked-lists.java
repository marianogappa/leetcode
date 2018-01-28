package main

import "fmt"

public class ListNode {
	int val;
	ListNode next;
	ListNode(int x) {
		val = x;
		next = null;
	}
}

public class Solution {
	public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
		int s1 = size(headA);
		int s2 = size(headB);
		if (s1 > s2) {
			int d = s1 - s2;
			int i;
			for (i = 1; i <= d; i++) {
				headA = headA.next;
			}
		}
		if (s2 > s1) {
			int d = s2 - s1;
			int i;
			for (i = 1; i <= d; i++) {
				headB = headB.next;
			}
		}
		while (headA != null) {
			if (headA == headB) {
				return headA;
			}
			headA = headA.next;
			headB = headB.next;
		}
		return null;
	}

	private int size(ListNode l) {
		if (l == null) {
			return 0;
		}
		return 1 + size(l.next);
	}
}
