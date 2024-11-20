# Created by Rafael Shen at 2024/11/20 10:11
# leetgo: 1.4.9
# https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-ii/

from typing import List

# @lc code=begin


class Solution:
    def shortestDistanceAfterQueries(
        self, n: int, queries: List[List[int]]
    ) -> List[int]:
        roads = [i + 1 for i in range(n)]
        res = []
        dist = n - 1
        for query in queries:
            k = roads[query[0]]
            roads[query[0]] = query[1]
            while k != -1 and k < query[1]:
                roads[k], k = -1, roads[k]
                dist -= 1
            res.append(dist)
        return res


# @lc code=end
