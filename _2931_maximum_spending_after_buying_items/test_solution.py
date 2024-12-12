import pytest
from .solution import Solution

test_cases = [
    pytest.param([[8, 5, 2], [6, 4, 1], [9, 7, 3]], 285, id="tc1"),
    pytest.param([[10, 8, 6, 4, 2], [9, 7, 5, 3, 2]], 386, id="tc2"),
]


@pytest.mark.parametrize("values, wanted", test_cases)
def test_maxSpending(values, wanted):
    assert Solution().maxSpending(values) == wanted
