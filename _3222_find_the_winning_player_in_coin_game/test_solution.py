import pytest
from .solution import Solution

test_cases = [
    pytest.param(2, 7, "Alice", id="tc1"),
    pytest.param(4, 11, "Bob", id="tc1"),
]


@pytest.mark.parametrize("x, y, wanted", test_cases)
def test_losingPlayer(x, y, wanted):
    assert Solution().losingPlayer(x, y) == wanted
