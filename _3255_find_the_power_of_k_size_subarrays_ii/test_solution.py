import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 2, 3, 4, 3, 2, 5], 3, [3, 4, -1, -1, -1], id="tc1"),
    pytest.param([2, 2, 2, 2, 2], 4, [-1, -1], id="tc2"),
    pytest.param([3, 2, 3, 2, 3, 2], 2, [-1, 3, -1, 3, -1], id="tc3"),
]


@pytest.mark.parametrize("nums, k, wanted", test_cases)
def test_resultsArray(nums, k, wanted):
    assert Solution().resultsArray(nums, k) == wanted
