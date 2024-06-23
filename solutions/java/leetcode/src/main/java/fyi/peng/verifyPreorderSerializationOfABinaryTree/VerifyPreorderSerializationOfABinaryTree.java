package fyi.peng.verifyPreorderSerializationOfABinaryTree;

import java.util.ArrayList;
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

}
