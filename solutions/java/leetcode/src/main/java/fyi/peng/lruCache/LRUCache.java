package fyi.peng.lruCache;

import fyi.peng.common.DoubleLinkList;
import fyi.peng.common.Node;

import java.util.HashMap;
import java.util.Map;

class LRUCache {

  private Map<Integer, Node<Pair<Integer, Integer>>> m;
  private DoubleLinkList<Pair<Integer, Integer>> l;
  int capacity;

  public LRUCache(int capacity) {
    this.capacity = capacity;
    m = new HashMap<>();
    l = new DoubleLinkList<>();
  }

  public int get(int key) {
    Node<Pair<Integer, Integer>> node = m.get(key);
    if (node == null) {
      return -1;
    }
    l.moveToFirst(node);
    return node.getData().v;
  }

  public void put(int key, int value) {
    Node<Pair<Integer, Integer>> node = m.get(key);
    if (node != null) {
      node.getData().v = value;
      l.moveToFirst(node);
      return;
    }

    if (m.size() == capacity) {
      Node<Pair<Integer, Integer>> tail = l.getTail();
      m.remove(tail.getData().k);
      l.remove(tail);
    }

    Node<Pair<Integer, Integer>> newNode = l.addFirst(new Pair<>(key, value));
    m.put(key, newNode);
  }

  static class Pair<K, V> {
    public K k;
    public V v;

    Pair(K k, V v) {
      this.k = k;
      this.v = v;
    }
  }
}
