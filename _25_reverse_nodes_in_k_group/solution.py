# Created by Rafael Shen at 2024/11/10 19:08
# leetgo: 1.4.9
# https://leetcode.cn/problems/reverse-nodes-in-k-group/

from typing import Optional
from leetgo_py import ListNode

# @lc code=begin


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        def reverse(reverse_head: ListNode):
            cur = reverse_head.next
            st = []
            while cur is not None:
                st.append(cur)
                cur = cur.next
            cur = reverse_head
            while len(st) > 0:
                cur.next = st.pop()
                cur = cur.next
            cur.next = None
            return cur

        tmp = ListNode(next=head)
        h, t = tmp, tmp.next
        cnt = 0
        while t is not None:
            cnt += 1
            if cnt == k:
                tnext = t.next
                t.next = None
                tail = reverse(h)
                tail.next = tnext
                cnt = 0
                h = tail
                t = tail.next
                continue
            t = t.next
        return tmp.next


# @lc code=end
