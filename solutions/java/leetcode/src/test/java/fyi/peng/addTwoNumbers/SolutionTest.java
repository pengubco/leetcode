package fyi.peng.addTwoNumbers;

import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.List;

import org.junit.jupiter.api.Test;

public class SolutionTest {
	@Test
	public void testAddTwoNumbers() {
		Solution s = new Solution();
		ListNode result = null;
		// 1 + 0
		result = s.addTwoNumbers(createList(List.of(1)), null);
		assertTrue(numberEqual(result, createList(List.of(1))));

		// 0 + 1
		result = s.addTwoNumbers(null, createList(List.of(1)));
		assertTrue(numberEqual(result, createList(List.of(1))));

		// 342 + 465 = 807
		result = s.addTwoNumbers(createList(List.of(2, 4, 3)), createList(List.of(5, 6, 4)));
		assertTrue(numberEqual(result, createList(List.of(7, 0, 8))));

		// 99 + 9 = 108
		result = s.addTwoNumbers(createList(List.of(9, 9)), createList(List.of(9)));
		assertTrue(numberEqual(result, createList(List.of(8, 0, 1))));
	}

	ListNode createList(List<Integer> num) {
		ListNode head = null;
		ListNode prev = null;
		for (int i : num) {
			ListNode cur = new ListNode(i);
			if (prev == null) {
				prev = cur;
				head = prev;
			} else {
				prev.next = cur;
				prev = cur;
			}
		}
		return head;
	}

	boolean numberEqual(ListNode l1, ListNode l2) {
		while (true) {
			if (l1 == null && l2==null) {
				return true;
			}
			if (l1 != null && l2 != null) {
				if (l1.val != l2.val) {
					return false;
				}
				l1 = l1.next;
				l2 = l2.next;
				continue;
			}
			return false;
		}
	}
}
