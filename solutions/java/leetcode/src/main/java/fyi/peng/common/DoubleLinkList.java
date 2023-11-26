package fyi.peng.common;

import java.util.ArrayList;
import java.util.List;

public class DoubleLinkList<T> {
	Node<T> head;
	Node<T> tail;

	public DoubleLinkList() {
	}

	public Node<T> addLast(T data) {
		Node<T> newNode = new Node<>(data, null, null);
		if (head == null) {
			head = tail = newNode;
			return newNode;
		}
		tail.next = newNode;
		newNode.prev = tail;
		tail = newNode;
		return newNode;
	}

	public Node<T> addFirst(T data) {
		Node<T> newNode = new Node<>(data, null, null);
		if (head == null) {
			head = tail = newNode;
			return newNode;
		}
		head.prev = newNode;
		newNode.next = head;
		head = newNode;
		return newNode;
	}

	public void moveToFirst(Node<T> node) {
		if (node == head) {
			return;
		}
		remove(node);
		head.prev = node;
		node.next = head;
		node.prev = null;
		head = node;
	}

	public void remove(Node<T> node) {
		Node<T> prev = node.prev;
		Node<T> next = node.next;
		if (prev == null) {
			head = next;
		} else {
			prev.next = next;
		}
		if (next == null) {
			tail = prev;
		} else {
			next.prev = prev;
		}
		node.prev = node.next = null;
		node = null;
	}

	public Node<T> last() {
		return tail;
	}

	public boolean isEmpty() {
		return head == null;
	}

	public List<T> toList() {
		List<T> result = new ArrayList<>();
		for (Node<T> cur = head; cur != null; cur = cur.next) {
			result.add(cur.data);
		}
		return result;
	}

	public Node<T> getHead() {
		return head;
	}

	public Node<T> getTail() {
		return tail;
	}

	@Override
	public String toString() {
		if (head == null) {
			return "";
		}
		StringBuilder sb = new StringBuilder();
		for (Node<T> cur = head; cur != null; cur = cur.next) {
			if (cur == head) {
				sb.append(String.format("%s", cur.data));
			} else {
				sb.append(String.format(" <-> %s ", cur.data));
			}
		}
		return sb.toString();
	}
}
