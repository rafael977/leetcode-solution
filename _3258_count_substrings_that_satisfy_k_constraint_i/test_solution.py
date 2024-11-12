import pytest
from .solution import Solution

test_cases = [
    pytest.param("10101", 1, 12, id="tc1"),
    pytest.param("1010101", 2, 25, id="tc2"),
    pytest.param("11111", 1, 15, id="tc3"),
]


@pytest.mark.parametrize("s, k, wanted", test_cases)
def test_countKConstraintSubstrings(s, k, wanted):
    assert Solution().countKConstraintSubstrings(s, k) == wanted
