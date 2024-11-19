# Created by Rafael Shen at 2024/11/19 22:24
# leetgo: 1.4.9
# https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/

from collections import deque
from typing import List

# @lc code=begin


class Solution:
    def shortestDistanceAfterQueries(
        self, n: int, queries: List[List[int]]
    ) -> List[int]:
        neighbors = [[i + 1] for i in range(n - 1)]
        neighbors.append([])
        ans = []
        for q in queries:
            neighbors[q[0]].append(q[1])
            ans.append(self.dfs(n, neighbors))
        return ans

    def dfs(self, n: int, neighbors: List[List[int]]) -> int:
        dist = [-1 for _ in range(n)]
        dist[0] = 0
        q = deque([0])
        while len(q) > 0:
            x = q.popleft()
            for y in neighbors[x]:
                if dist[y] >= 0:
                    continue
                dist[y] = dist[x] + 1
                q.append(y)
        return dist[n - 1]


# @lc code=end
