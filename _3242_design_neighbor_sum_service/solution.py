# Created by Rafael Shen at 2024/11/09 11:36
# leetgo: 1.4.9
# https://leetcode.cn/problems/design-neighbor-sum-service/

from typing import List

# @lc code=begin


class NeighborSum:

    def __init__(self, grid: List[List[int]]):
        self.m = len(grid)
        self.n = len(grid[0])
        self.grid = grid
        self.dict = dict()
        for i in range(self.m):
            for j in range(self.n):
                self.dict[grid[i][j]] = (i, j)

    def adjacentSum(self, value: int) -> int:
        dirs = [(0, -1), (1, 0), (0, 1), (-1, 0)]
        i, j = self.dict[value]
        neighbour_sum = 0
        for d in dirs:
            ii, jj = i + d[0], j + d[1]
            if ii >= 0 and ii < self.m and jj >= 0 and jj < self.n:
                neighbour_sum += self.grid[ii][jj]
        return neighbour_sum

    def diagonalSum(self, value: int) -> int:
        dirs = [(-1, -1), (-1, 1), (1, -1), (1, 1)]
        i, j = self.dict[value]
        neighbour_sum = 0
        for d in dirs:
            ii, jj = i + d[0], j + d[1]
            if ii >= 0 and ii < self.m and jj >= 0 and jj < self.n:
                neighbour_sum += self.grid[ii][jj]
        return neighbour_sum


# Your NeighborSum object will be instantiated and called as such:
# obj = NeighborSum(grid)
# param_1 = obj.adjacentSum(value)
# param_2 = obj.diagonalSum(value)

# @lc code=end
