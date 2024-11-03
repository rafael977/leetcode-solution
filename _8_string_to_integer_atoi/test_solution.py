import pytest
from .solution import Solution

test_cases = [
    pytest.param("42", 42, id="tc1"),
    pytest.param(" -42", -42, id="tc2"),
    pytest.param("1337c0d3", 1337, id="tc3"),
    pytest.param("0-1", 0, id="tc4"),
    pytest.param("words and 987", 0, id="tc5"),
    pytest.param("-91283472332", -2147483648, id="tc6"),
    pytest.param("  -0012a42", -12, id="tc7"),
]


@pytest.mark.parametrize("s, wanted", test_cases)
def test_(s, wanted):
    assert Solution().myAtoi(s) == wanted
