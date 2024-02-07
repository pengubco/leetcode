package fyi.peng.decodeString;

// https://leetcode.com/problems/decode-string
public class DecodeString {
	public String decodeString(String s) {
		StringBuilder sb = new StringBuilder();
		while (s.length()>0) {
			var start = s.indexOf('[');
			if (start == -1) {
				sb.append(s);
				return sb.toString();
			}
			var digitIndex = indexOfDigit(s);
			if (digitIndex == -1) {
				sb.append(s);
				return sb.toString();
			}
			sb.append(s.substring(0, digitIndex));
			var end = findClosingBracket(s.substring(start+1));
			if (end==-1) {
				// should never happen.
				return "";
			}
			end += start + 1;
			var stringToRepeat = decodeString(s.substring(start+1, end));
			var repeat = Integer.parseInt(s.substring(digitIndex, start));
			for (int j=0; j<repeat; j++) {
				sb.append(stringToRepeat);
			}
			s = s.substring(end+1);
		}
		return sb.toString();
	}

	public int indexOfDigit(String s) {
		for (int i=0; i<s.length(); i++ ) {
			if (s.charAt(i) >= '0' && s.charAt(i)<='9') {
				return i;
			}
		}
		return -1;
	}

	public int findClosingBracket(String s) {
		int i = 0;
		int openCount = 0;
		while (i<s.length()) {
			char c = s.charAt(i);
			if (c != '[' && c!=']') {
				i++;
			} else if (c=='[') {
				openCount++;
				i++;
			} else if (openCount == 0) {
				return i;
			} else {
				openCount--;
				i++;
			}
		}
		return -1;
	}
}
