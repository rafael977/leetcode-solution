import pytest
from .solution import Solution

test_cases = [
    pytest.param("III", 3, id="tc1"),
    pytest.param("LVIII", 58, id="tc2"),
    pytest.param("MCMXCIV", 1994, id="tc3"),
]


@pytest.mark.parametrize("s, wanted", test_cases)
def test_romanToInt(s, wanted):
    assert Solution().romanToInt(s) == wanted
