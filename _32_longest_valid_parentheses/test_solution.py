import pytest
from .solution import Solution

test_cases = [
    pytest.param("(()", 2, id="tc1"),
    pytest.param(")()())", 4, id="tc2"),
    pytest.param("", 0, id="tc3"),
]


@pytest.mark.parametrize("s, wanted", test_cases)
def test_longestValidParentheses(s, wanted):
    assert Solution().longestValidParentheses(s) == wanted
