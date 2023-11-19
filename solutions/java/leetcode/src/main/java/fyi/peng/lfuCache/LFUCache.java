package fyi.peng.lfuCache;

import fyi.peng.common.DoubleLinkList;
import fyi.peng.common.Node;

import java.util.HashMap;
import java.util.Map;
import java.util.TreeMap;

public class LFUCache {

  Map<Integer, Node<Data>> m;
  TreeMap<Integer, DoubleLinkList<Data>> dlls;

  int capacity;

  public LFUCache(int capacity) {
    this.capacity = capacity;
    this.m = new HashMap<>();
    this.dlls = new TreeMap<>();
  }

  public int get(int key) {
    Node<Data> node = m.get(key);
    if (node == null) {
      return -1;
    }
    int result = node.getData().value;
    increaseFrequency(node);
    return result;
  }

  public void put(int key, int value) {
    Node<Data> node = m.get(key);
    if (node != null) {
      node.getData().value = value;
      increaseFrequency(node);
      return;
    }

    if (m.size() == capacity) {
      Map.Entry<Integer, DoubleLinkList<Data>> entry = dlls.firstEntry();
      DoubleLinkList<Data> dll = entry.getValue();
      Node<Data> nodeToDelete = dll.last();
      m.remove(nodeToDelete.getData().key);
      dll.remove(nodeToDelete);
    }

    DoubleLinkList<Data> dll = dlls.get(1);
    if (dll == null) {
      dll = new DoubleLinkList<>();
      dlls.put(1, dll);
    }
    Node<Data> newNode = dll.addFirst(new Data(key, value, 1));
    m.put(key, newNode);
  }

  public void increaseFrequency(Node<Data> node) {
    // remove from old frequency list.
    int freq = node.getData().frequency;
    DoubleLinkList<Data> oldDll = dlls.get(freq);
    oldDll.remove(node);
    if (oldDll.isEmpty()) {
      dlls.remove(freq);
    }

    // add to new frequency list.
    DoubleLinkList<Data> newDll = dlls.get(freq + 1);
    if (newDll == null) {
      newDll = new DoubleLinkList<>();
      dlls.put(freq+1, newDll);
    }
    Node<Data> newNode = newDll.addFirst(node.getData());
    newNode.getData().frequency++;
    m.put(node.getData().key, newNode);
  }

  static class Data {
    public int key;
    public int value;
    public int frequency;

    Data(int key, int value, int frequency) {
      this.key = key;
      this.value = value;
      this.frequency = frequency;
    }
  }
}