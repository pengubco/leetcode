package fyi.peng.common;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class DoubleLinkListTest {

  @Test
  public void testSimpleAddRemove() {
    DoubleLinkList<Integer> l = new DoubleLinkList<>();
    l.addLast(1);
    l.addFirst(10);
    l.addFirst(11);
    Assertions.assertIterableEquals(Arrays.asList(11, 10, 1), l.toList());

    l.moveToFirst(l.tail);
    Assertions.assertIterableEquals(Arrays.asList(1, 11, 10), l.toList());

    l.remove(l.head);
    l.remove(l.head);
    l.remove(l.tail);
    Assertions.assertIterableEquals(Arrays.asList(), l.toList());
  }

}
