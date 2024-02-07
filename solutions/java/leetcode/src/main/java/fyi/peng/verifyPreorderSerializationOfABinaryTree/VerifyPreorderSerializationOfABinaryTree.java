package fyi.peng.verifyPreorderSerializationOfABinaryTree;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

// https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree
public class VerifyPreorderSerializationOfABinaryTree {
	public boolean isValidSerialization(String preorder) {
		String[] split = preorder.split(",");
		List<String> stack = new ArrayList<>();
		for (String s : split) {
			removeLeaf(stack);
			stack.add(s);
		}
		removeLeaf(stack);

		return stack.size() == 1 && stack.get(0).equals("#");
	}


	public void removeLeaf(List<String> stack) {
		while(true) {
			int n = stack.size();
			if (n >= 3 && stack.get(n-1).equals("#") && stack.get(n-2).equals("#") && !stack.get(n-3).equals("#")) {
				stack.remove(n-3);
				stack.remove(n-3);
				continue;
			}
			break;
		}
	}

	// slower version that not use stack. 
	public boolean isValidSerialization1(String preorder) {
		LinkedList<String> ll = new LinkedList<>();
		String[] split = preorder.split(",");
		for (String s : split) {
			ll.add(s);
		}

		// remove leaf iteratively.
		int startIndex = 0;
		while (ll.size() > 1) {
			var leafIndex = -1;
			for (int i=startIndex; i+2 < ll.size(); i++) {
				if (!ll.get(i).equals("#") && ll.get(i+1).equals("#") && ll.get(i+2).equals("#")) {
					leafIndex = i;
					break;
				}
			}
			if (leafIndex == -1) {
				return false;
			}
			// remove the leaf node and its two children. 
			ll.remove(leafIndex);
			ll.remove(leafIndex);
			// no need to start from 0. 
			startIndex = Math.max(0, leafIndex - 2);
		}
		return ll.size() == 1 && ll.get(0).equals("#");
	}
}
