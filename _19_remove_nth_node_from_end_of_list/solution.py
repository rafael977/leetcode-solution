# Created by Rafael Shen at 2024/11/08 11:56
# leetgo: 1.4.9
# https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

from typing import *
from leetgo_py import *

# @lc code=begin


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        b = head
        for i in range(n):
            b = b.next
        tmp = ListNode(next=head)
        a = tmp
        while b is not None:
            b = b.next
            a = a.next
        a.next = a.next.next
        return tmp.next


# @lc code=end
