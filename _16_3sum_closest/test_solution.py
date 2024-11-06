import pytest
from .solution import Solution

test_cases = [
    pytest.param([-1, 2, 1, -4], 1, 2, id="tc1"),
    pytest.param([0, 0, 0], 1, 0, id="tc2"),
    pytest.param([1, 1, 1, 1], 0, 3, id="tc3"),
    pytest.param([1, 1, 1, 0], -100, 2, id="tc4"),
]


@pytest.mark.parametrize("nums, target, wanted", test_cases)
def test_threeSumClosest(nums, target, wanted):
    assert Solution().threeSumClosest(nums, target) == wanted
