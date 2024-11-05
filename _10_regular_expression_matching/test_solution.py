import pytest
from .solution import Solution

test_cases = [
    pytest.param("aa", "a", False, id="tc1"),
    pytest.param("aa", "a*", True, id="tc2"),
    pytest.param("ab", ".*", True, id="tc3"),
]


@pytest.mark.parametrize("s, p, wanted", test_cases)
def test_isMatch(s, p, wanted):
    assert Solution().isMatch(s, p) == wanted
