import pytest
from .solution import Solution

test_cases = [
    pytest.param(3749, "MMMDCCXLIX", id="tc1"),
    pytest.param(58, "LVIII", id="tc2"),
    pytest.param(1994, "MCMXCIV", id="tc3"),
]


@pytest.mark.parametrize("num, wanted", test_cases)
def test_intToRoman(num, wanted):
    assert Solution().intToRoman(num) == wanted
