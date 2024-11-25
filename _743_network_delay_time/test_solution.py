import pytest
from .solution import Solution

test_cases = [
    pytest.param([[2, 1, 1], [2, 3, 1], [3, 4, 1]], 4, 2, 2, id="tc1"),
    pytest.param([[1, 2, 1]], 2, 1, 1, id="tc2"),
    pytest.param([[1, 2, 1]], 2, 2, -1, id="tc3"),
    pytest.param([[1, 2, 1], [2, 3, 2], [1, 3, 4]], 3, 1, 3, id="tc4"),
]


@pytest.mark.parametrize("times, n, k, wanted", test_cases)
def test_networkDelayTime(times, n, k, wanted):
    assert Solution().networkDelayTime(times, n, k) == wanted
