package fyi.peng.common;

public class Node<T> {
  T data;
  Node<T> prev;
  Node<T> next;

  public Node(T data, Node<T> prev, Node<T> next) {
    this.data = data;
    this.prev = prev;
    this.next = next;
  }

  public T getData() {
    return data;
  }

  public void setData(T data) {
    this.data = data;
  }
}
