# Created by Rafael Shen at 2024/11/25 10:03
# leetgo: 1.4.11
# https://leetcode.cn/problems/network-delay-time/

from typing import List
import heapq

# @lc code=begin


class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        links = [[] for _ in range(n)]
        for t in times:
            links[t[0] - 1].append((t[1] - 1, t[2]))

        visited = [False] * n
        q = []
        heapq.heappush(q, (0, k - 1))
        dt = 0
        while len(q) > 0:
            cw, u = heapq.heappop(q)
            if visited[u]:
                continue
            visited[u] = True
            dt = max(dt, cw)
            for nx in links[u]:
                if visited[nx[0]]:
                    continue
                w = cw + nx[1]
                heapq.heappush(q, (w, nx[0]))
        return -1 if not all(visited) else dt


# @lc code=end
