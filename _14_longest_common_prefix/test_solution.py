import pytest
from .solution import Solution

test_cases = [
    pytest.param(["flower", "flow", "flight"], "fl", id="tc1"),
    pytest.param(["dog", "racecar", "car"], "", id="tc2"),
]


@pytest.mark.parametrize("strs, wanted", test_cases)
def test_longestCommonPrefix(strs, wanted):
    assert Solution().longestCommonPrefix(strs) == wanted
