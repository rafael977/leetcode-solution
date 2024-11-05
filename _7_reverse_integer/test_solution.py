import pytest
from .solution import Solution

test_cases = [
    pytest.param(123, 321, id="tc1"),
    pytest.param(-123, -321, id="tc2"),
    pytest.param(120, 21, id="tc3"),
    pytest.param(-1563847412, 0, id="tc4"),
]


@pytest.mark.parametrize("x, wanted", test_cases)
def test_reverse(x, wanted):
    assert Solution().reverse(x) == wanted
