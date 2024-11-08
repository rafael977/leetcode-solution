import pytest
from .solution import Solution

test_cases = [
    pytest.param(3, ["((()))", "(()())", "(())()", "()(())", "()()()"], id="tc1"),
    pytest.param(1, ["()"], id="tc2"),
]


@pytest.mark.parametrize("n, wanted", test_cases)
def test_generateParenthesis(n, wanted):
    assert Solution().generateParenthesis(n) == wanted
