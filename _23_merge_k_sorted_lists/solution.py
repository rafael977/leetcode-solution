# Created by Rafael Shen at 2024/11/08 14:30
# leetgo: 1.4.9
# https://leetcode.cn/problems/merge-k-sorted-lists/

from typing import *
from leetgo_py import *

# @lc code=begin

from queue import PriorityQueue


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        head = ListNode()
        cur = head
        q = PriorityQueue()
        for i in range(len(lists)):
            if lists[i]:
                q.put((lists[i].val, i))
        while not q.empty():
            v, i = q.get()
            cur.next = ListNode(v)
            cur = cur.next
            if lists[i].next:
                lists[i] = lists[i].next
                q.put((lists[i].val, i))
        return head.next


# @lc code=end

if __name__ == "__main__":
    lists: List[ListNode] = deserialize("List[ListNode]", read_line())
    ans = Solution().mergeKLists(lists)
    print("\noutput:", serialize(ans, "ListNode"))
