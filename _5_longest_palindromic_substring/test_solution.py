import pytest
from .solution import Solution

test_cases = [
    pytest.param("babad", "bab", id="tc1"),
    pytest.param("cbbd", "bb", id="tc2"),
]


@pytest.mark.parametrize("s, wanted", test_cases)
def test_longestPalindrome(s, wanted):
    assert Solution().longestPalindrome(s) == wanted
