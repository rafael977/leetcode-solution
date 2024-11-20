import pytest
from .solution import Solution

test_cases = [
    pytest.param(5, [[2, 4], [0, 2], [0, 4]], [3, 2, 1], id="tc1"),
    pytest.param(4, [[0, 3], [0, 2]], [1, 1], id="tc2"),
]


@pytest.mark.parametrize("n, queries, wanted", test_cases)
def test_shortestDistanceAfterQueries(n, queries, wanted):
    assert Solution().shortestDistanceAfterQueries(n, queries) == wanted
