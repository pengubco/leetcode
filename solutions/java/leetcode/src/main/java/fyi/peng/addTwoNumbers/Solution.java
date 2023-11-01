package fyi.peng.addTwoNumbers;

public class Solution {
	public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
		ListNode head = null;
		ListNode prev = null;
		int carry = 0;
		while (l1 != null || l2 != null || carry > 0) {
			int sum = carry;
			if (l1 != null) {
				sum += l1.val;
			}
			if (l2 != null) {
				sum += l2.val;
			}
			if (sum >= 10) {
				sum -= 10;
				carry = 1;
			} else {
				carry = 0;
			}
			ListNode cur = new ListNode(sum);
			if (prev == null) {
				prev = cur;
				head = cur;
			} else {
				prev.next = cur;
				prev = cur;
			}
			if (l1 != null) {
				l1 = l1.next;
			}
			if (l2 != null) {
				l2 = l2.next;
			}
		}
		return head;
	}
}