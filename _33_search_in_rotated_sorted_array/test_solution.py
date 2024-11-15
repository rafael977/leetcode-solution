import pytest
from .solution import Solution

test_cases = [
    pytest.param([4, 5, 6, 7, 0, 1, 2], 0, 4, id="tc1"),
    pytest.param([4, 5, 6, 7, 0, 1, 2], 3, -1, id="tc2"),
    pytest.param([1], 0, -1, id="tc3"),
    pytest.param([3, 1], 1, 1, id="tc4"),
]


@pytest.mark.parametrize("nums, target, wanted", test_cases)
def test_search(nums, target, wanted):
    assert Solution().search(nums, target) == wanted
