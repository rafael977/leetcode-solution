import pytest
from .solution import Solution

test_cases = [
    pytest.param(
        [10, 1, 2, 7, 6, 1, 5], 8, [[1, 1, 6], [1, 2, 5], [1, 7], [2, 6]], id="tc1"
    ),
    pytest.param([2, 5, 2, 1, 2], 5, [[1, 2, 2], [5]], id="tc2"),
]


@pytest.mark.parametrize("candidates, target, wanted", test_cases)
def test_combinationSum2(candidates, target, wanted):
    assert Solution().combinationSum2(candidates, target) == wanted
