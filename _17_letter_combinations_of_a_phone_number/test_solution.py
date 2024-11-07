import pytest
from .solution import Solution

test_cases = [
    pytest.param(
        "23", ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"], id="tc1"
    ),
    pytest.param("", [], id="tc2"),
    pytest.param("2", ["a", "b", "c"], id="tc3"),
]


@pytest.mark.parametrize("digits, wanted", test_cases)
def test_letterCombinations(digits, wanted):
    assert Solution().letterCombinations(digits) == wanted
