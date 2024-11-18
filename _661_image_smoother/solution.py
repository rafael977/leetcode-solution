# Created by Rafael Shen at 2024/11/18 09:50
# leetgo: 1.4.9
# https://leetcode.cn/problems/image-smoother/

from typing import List

# @lc code=begin


class Solution:
    def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
        m, n = len(img), len(img[0])
        ans = [[0] * n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                sum = 0
                cnt = 0
                for k in range(-1, 2):
                    for l in range(-1, 2):
                        if 0 <= i + k < m and 0 <= j + l < n:
                            sum += img[i + k][j + l]
                            cnt += 1
                ans[i][j] = sum // cnt
        return ans


# @lc code=end
