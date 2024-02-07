package fyi.peng.decodeString;


import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class DecodeStringTest {

	@Test
	public void testDecodeString() {
		DecodeString sol = new DecodeString();

		assertEquals("aaabcbc", sol.decodeString("3[a]2[bc]"));
		assertEquals("accaccacc", sol.decodeString("3[a2[c]]"));
		assertEquals("abcabccdcdcdef", sol.decodeString("2[abc]3[cd]ef"));
	}
}
