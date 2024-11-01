from .solution import Solution
import pytest

test_cases = [
    pytest.param([2, 7, 11, 15], 9, [0, 1], id="tc1"),
    pytest.param([3, 2, 4], 6, [1, 2], id="tc2"),
    pytest.param([3, 3], 6, [0, 1], id="tc3"),
]


@pytest.mark.parametrize("nums, target, wanted", test_cases)
def test_twoSum(nums, target, wanted):
    assert Solution().twoSum(nums, target) == wanted
