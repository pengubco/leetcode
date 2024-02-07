package fyi.peng.verifyPreorderSerializationOfABinaryTree;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class VerifyPreorderSerializationOfABinaryTreeTest {
	
	@Test
	public void testPreorder() {
		VerifyPreorderSerializationOfABinaryTree sol = new VerifyPreorderSerializationOfABinaryTree();

		assertEquals(true, sol.isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"));
		assertEquals(false, sol.isValidSerialization("1,#"));
		assertEquals(false, sol.isValidSerialization("9,#,#,1"));
	}
}
