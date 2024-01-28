package fyi.peng.minimumSizeSubarrayInInfiniteArray;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class SolutionTest {

    @ParameterizedTest
    @MethodSource
    public void testMinSizeSubarray(int[] nums, int target, int expected) {
        Solution s = new Solution();
        int got = s.minSizeSubarray(nums, target);
        assertEquals(expected, got);
    }

    private static Stream<Arguments> testMinSizeSubarray() {
        return Stream.of(
                Arguments.of(new int[] { 1, 2, 3 }, 5, 2),
                Arguments.of(new int[] { 1, 1, 1, 2, 3 }, 4, 2),
                Arguments.of(new int[] { 2, 4, 6, 8 }, 3, -1),
                Arguments.of(new int[] { 10, 10, 1, 1, 1, 1 }, 40, -1));
    }
}
