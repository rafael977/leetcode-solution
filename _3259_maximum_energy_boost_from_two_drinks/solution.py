# Created by Rafael Shen at 2024/11/01 12:18
# leetgo: 1.4.9
# https://leetcode.cn/problems/maximum-energy-boost-from-two-drinks/

from typing import *
from leetgo_py import *

# @lc code=begin


class Solution:
    def maxEnergyBoost(self, energyDrinkA: List[int], energyDrinkB: List[int]) -> int:
        a, b = energyDrinkA[0], energyDrinkB[0]
        for i in range(1, len(energyDrinkA)):
            a, b = max(a + energyDrinkA[i], b), max(a, b + energyDrinkB[i])
        return max(a, b)


# @lc code=end
