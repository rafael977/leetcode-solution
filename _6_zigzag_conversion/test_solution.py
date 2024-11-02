import pytest
from .solution import Solution

test_cases = [
    pytest.param("PAYPALISHIRING", 3, "PAHNAPLSIIGYIR", id="tc1"),
    pytest.param("PAYPALISHIRING", 4, "PINALSIGYAHRPI", id="tc2"),
    pytest.param("A", 1, "A", id="tc3"),
    pytest.param("ABC", 1, "ABC", id="tc3"),
]


@pytest.mark.parametrize("s, numRows, wanted", test_cases)
def test_(s, numRows, wanted):
    assert Solution().convert(s, numRows) == wanted
