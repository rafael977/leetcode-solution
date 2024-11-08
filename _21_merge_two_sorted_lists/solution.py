# Created by Rafael Shen at 2024/11/08 12:25
# leetgo: 1.4.9
# https://leetcode.cn/problems/merge-two-sorted-lists/

from typing import *
from leetgo_py import *

# @lc code=begin


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
        tmp = ListNode()
        cur = tmp
        while list1 is not None and list2 is not None:
            if list1.val <= list2.val:
                cur.next = list1
                list1 = list1.next
            else:
                cur.next = list2
                list2 = list2.next
            cur = cur.next
        if list1 is not None:
            cur.next = list1
        if list2 is not None:
            cur.next = list2

        return tmp.next


# @lc code=end
