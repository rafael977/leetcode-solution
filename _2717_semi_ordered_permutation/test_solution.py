import pytest
from .solution import Solution

test_cases = [
    pytest.param([2, 1, 4, 3], 2, id="tc1"),
    pytest.param([2, 4, 1, 3], 3, id="tc2"),
    pytest.param([1, 3, 4, 2, 5], 0, id="tc3"),
]


@pytest.mark.parametrize("nums, wanted", test_cases)
def test_semiOrderedPermutation(nums, wanted):
    assert Solution().semiOrderedPermutation(nums) == wanted
