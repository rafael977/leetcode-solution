# Created by Rafael Shen at 2024/11/03 13:20
# leetgo: 1.4.9
# https://leetcode.cn/problems/shopping-offers/

from functools import lru_cache
from typing import *

# @lc code=begin


class Solution:
    def shoppingOffers(
        self, price: List[int], special: List[List[int]], needs: List[int]
    ) -> int:
        n = len(needs)

        @lru_cache(None)
        def dfs(cur_needs):
            min_price = sum(price[i] * need for i, need in enumerate(cur_needs))
            for cur_special in special:
                special_price = cur_special[-1]
                next_needs = []
                for i in range(n):
                    if cur_special[i] > cur_needs[i]:
                        break
                    next_needs.append(cur_needs[i] - cur_special[i])
                if len(next_needs) == n:
                    min_price = min(min_price, special_price + dfs(tuple(next_needs)))
            return min_price

        return dfs(tuple(needs))


# @lc code=end
