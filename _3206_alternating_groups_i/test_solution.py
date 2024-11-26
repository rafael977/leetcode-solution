import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 1, 1], 0, id="tc1"),
    pytest.param([0, 1, 0, 0, 1], 3, id="tc2"),
]


@pytest.mark.parametrize("colors, wanted", test_cases)
def test_numberOfAlternatingGroups(colors, wanted):
    assert Solution().numberOfAlternatingGroups(colors) == wanted
