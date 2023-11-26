
1. Like the LRU, we use a double linked list(DLL) to move recent node to HEAD and evict the TAIL.
2. Unlike the LRU, we move the recent node to the virtual HEAD where all nodes of the frequency begins.
3. To simplify coding, we create a DLL for each frequency. This makes it much easier to read 
   than keeping all nodes in one DLL and keep track of HEADs for each frequency.
4. On get/put: move a node from one DLL to the head of another DLL.
5. There is no DLL structure that supports constant add or remove of any node by reference. So we 
	 need to implement DLL on our own.
6. Use a TreeMap to map frequency to DLL. TreeMap keeps frequencies sorted. However, 
   the time complexity would be O(log_N). It passes the online check though. 
