import pytest
from .solution import Solution

tests_cases = [
    pytest.param([5, 7, 7, 8, 8, 10], 8, [3, 4], id="tc1"),
    pytest.param([5, 7, 7, 8, 8, 10], 6, [-1, -1], id="tc2"),
    pytest.param([], 0, [-1, -1], id="tc3"),
    pytest.param([1], 1, [0, 0], id="tc4"),
]


@pytest.mark.parametrize("nums, target, wanted", tests_cases)
def test_searchRange(nums, target, wanted):
    assert Solution().searchRange(nums, target) == wanted
