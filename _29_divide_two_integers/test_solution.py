import pytest
from .solution import Solution

test_cases = [
    pytest.param(10, 3, 3, id="tc1"),
    pytest.param(7, -3, -2, id="tc2"),
    pytest.param(-2147483648, -1, 2147483647, id="tc3"),
]


@pytest.mark.parametrize("dividend, divisor, wanted", test_cases)
def test_divide(dividend, divisor, wanted):
    assert Solution().divide(dividend, divisor) == wanted
