import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 2, 3], [1, 3, 2], id="tc1"),
    pytest.param([3, 2, 1], [1, 2, 3], id="tc2"),
    pytest.param([1, 1, 5], [1, 5, 1], id="tc3"),
    pytest.param([1, 3, 2], [2, 1, 3], id="tc4"),
    pytest.param([2, 3, 1], [3, 1, 2], id="tc5"),
]


@pytest.mark.parametrize("nums, wanted", test_cases)
def test_nextPermutation(nums, wanted):
    assert Solution().nextPermutation(nums) == wanted
