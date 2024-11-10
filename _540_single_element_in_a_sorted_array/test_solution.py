import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 1, 2, 3, 3, 4, 4, 8, 8], 2, id="tc1"),
    pytest.param([3, 3, 7, 7, 10, 11, 11], 10, id="tc2"),
]


@pytest.mark.parametrize("nums, wanted", test_cases)
def test_singleNonDuplicate(nums, wanted):
    assert Solution().singleNonDuplicate(nums) == wanted
