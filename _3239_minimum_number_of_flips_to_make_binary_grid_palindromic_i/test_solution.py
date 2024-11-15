import pytest
from .solution import Solution

test_cases = [
    pytest.param([[1, 0, 0], [0, 0, 0], [0, 0, 1]], 2, id="tc1"),
    pytest.param([[0, 1], [0, 1], [0, 0]], 1, id="tc2"),
    pytest.param([[1], [0]], 0, id="tc3"),
]


@pytest.mark.parametrize("grid, wanted", test_cases)
def test_(grid, wanted):
    assert Solution().minFlips(grid) == wanted
