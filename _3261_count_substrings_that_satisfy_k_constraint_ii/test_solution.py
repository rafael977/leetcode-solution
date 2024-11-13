import pytest
from .solution import Solution

test_cases = [
    pytest.param("0001111", 2, [[0, 6]], [26], id="tc1"),
    pytest.param("010101", 1, [[0, 5], [1, 4], [2, 3]], [15, 9, 3], id="tc2"),
]


@pytest.mark.parametrize("s, k, queries, wanted", test_cases)
def test_countKConstraintSubstrings(s, k, queries, wanted):
    assert Solution().countKConstraintSubstrings(s, k, queries) == wanted
