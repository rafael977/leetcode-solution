# Created by Rafael Shen at 2024/11/13 15:08
# leetgo: 1.4.9
# https://leetcode.cn/problems/longest-valid-parentheses/

# @lc code=begin


class Solution:
    def longestValidParentheses(self, s: str) -> int:
        st = [-1]
        ans = 0
        for i, c in enumerate(s):
            if c == "(":
                st.append(i)
            else:
                st.pop()
                if len(st) > 0:
                    ans = max(ans, i - st[-1])
                else:
                    st.append(i)

        return ans


# @lc code=end
