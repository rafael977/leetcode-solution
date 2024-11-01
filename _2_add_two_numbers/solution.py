# Created by Rafael Shen at 2024/11/01 16:34
# leetgo: 1.4.9
# https://leetcode.cn/problems/add-two-numbers/

from typing import *
from leetgo_py import *

# @lc code=begin


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        ans = ListNode()
        cur = ans
        carry = 0
        while l1 is not None or l2 is not None:
            v = carry
            if l1 is not None:
                v += l1.val
                l1 = l1.next
            if l2 is not None:
                v += l2.val
                l2 = l2.next
            cur.next = ListNode(v % 10)
            cur = cur.next
            carry = v // 10
        if carry != 0:
            cur.next = ListNode(carry)

        return ans.next


# @lc code=end
