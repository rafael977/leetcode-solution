import pytest
from .solution import Solution

test_cases = [
    pytest.param([2, 1, 3, 5, 6], 5, 2, [8, 4, 6, 5, 6], id="tc1"),
    pytest.param([1, 2], 3, 4, [16, 8], id="tc2"),
]


@pytest.mark.parametrize("nums, k, multiplier, wanted", test_cases)
def test_getFinalState(nums, k, multiplier, wanted):
    assert Solution().getFinalState(nums, k, multiplier) == wanted
