import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 3, 5, 6], 5, 2, id="tc1"),
    pytest.param([1, 3, 5, 6], 2, 1, id="tc2"),
    pytest.param([1, 3, 5, 6], 7, 4, id="tc3"),
]


@pytest.mark.parametrize("nums, target, wanted", test_cases)
def test_searchInsert(nums, target, wanted):
    assert Solution().searchInsert(nums, target) == wanted
