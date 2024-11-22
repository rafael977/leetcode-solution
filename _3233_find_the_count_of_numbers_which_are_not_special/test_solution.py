import pytest
from .solution import Solution

test_cases = [
    pytest.param(5, 7, 3, id="tc1"),
    pytest.param(4, 16, 11, id="tc2"),
    pytest.param(1, 2, 2, id="tc3"),
    pytest.param(182, 18677, 18470, id="tc4"),
]


@pytest.mark.parametrize("l, r, wanted", test_cases)
def test_nonSpecialCount(l, r, wanted):
    assert Solution().nonSpecialCount(l, r) == wanted
