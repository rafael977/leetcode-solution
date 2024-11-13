import pytest
from .solution import Solution

test_cases = [
    pytest.param(7, [1, 3, 4, 5], 16, id="tc1"),
    pytest.param(9, [5, 6, 1, 4, 2], 22, id="tc2"),
]


@pytest.mark.parametrize("n, cuts, wanted", test_cases)
def test_minCost(n, cuts, wanted):
    assert Solution().minCost(n, cuts) == wanted
