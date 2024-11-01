import pytest
from .solution import Solution

test_cases = [
    pytest.param("abcabcbb", 3, id="tc1"),
    pytest.param("bbbbb", 1, id="tc2"),
    pytest.param("pwwkew", 3, id="tc3"),
]


@pytest.mark.parametrize("s, wanted", test_cases)
def test_lengthOfLongestSubstring(s, wanted):
    assert Solution().lengthOfLongestSubstring(s) == wanted
