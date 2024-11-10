# Created by Rafael Shen at 2024/11/10 13:34
# leetgo: 1.4.9
# https://leetcode.cn/problems/swap-nodes-in-pairs/

from typing import *
from leetgo_py import *

# @lc code=begin


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head
        tmp = ListNode(next=head)
        prev = tmp
        n1, n2 = head, head.next
        while n1 is not None and n2 is not None:
            prev.next = n2
            n1.next = n2.next
            n2.next = n1
            prev = n1
            if not n1.next or not n1.next.next:
                break
            n2 = n1.next.next
            n1 = n1.next
        return tmp.next


# @lc code=end
