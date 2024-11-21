import pytest
from .solution import Solution

test_cases = [
    pytest.param([2, 3, 6, 7], 7, [[2, 2, 3], [7]], id="tc1"),
    pytest.param([2, 3, 5], 8, [[2, 2, 2, 2], [2, 3, 3], [3, 5]], id="tc2"),
    pytest.param([2], 1, [], id="tc3"),
]


@pytest.mark.parametrize("candidates, target, wanted", test_cases)
def test_combinationSum(candidates, target, wanted):
    assert Solution().combinationSum(candidates, target) == wanted
