import pytest
from .solution import Solution

test_cases = [
    pytest.param(
        [1, 0, -1, 0, -2, 2],
        0,
        [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]],
        id="tc1",
    ),
    pytest.param(
        [2, 2, 2, 2, 2],
        8,
        [[2, 2, 2, 2]],
        id="tc2",
    ),
    pytest.param(
        [0, 0, 0, 0],
        0,
        [[0, 0, 0, 0]],
        id="tc3",
    ),
    pytest.param(
        [1, -2, -5, -4, -3, 3, 3, 5],
        -11,
        [[-5, -4, -3, 1]],
        id="tc4",
    ),
]


@pytest.mark.parametrize("nums, target, wanted", test_cases)
def test_fourSum(nums, target, wanted):
    assert sorted(Solution().fourSum(nums, target)) == sorted(wanted)
