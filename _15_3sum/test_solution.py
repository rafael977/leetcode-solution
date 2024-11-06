import pytest
from .solution import Solution

test_cases = [
    pytest.param([-1, 0, 1, 2, -1, -4], [[-1, -1, 2], [-1, 0, 1]], id="tc1"),
    pytest.param([0, 1, 1], [], id="tc2"),
    pytest.param([0, 0, 0], [[0, 0, 0]], id="tc3"),
]


@pytest.mark.parametrize("nums, wanted", test_cases)
def test_threeSum(nums, wanted):
    assert sorted(Solution().threeSum(nums)) == sorted(wanted)
