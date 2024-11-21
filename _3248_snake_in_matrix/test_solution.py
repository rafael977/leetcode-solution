import pytest
from .solution import Solution

test_cases = [
    pytest.param(2, ["RIGHT", "DOWN"], 3, id="tc1"),
    pytest.param(3, ["DOWN", "RIGHT", "UP"], 1, id="tc2"),
]


@pytest.mark.parametrize("n, commands, wanted", test_cases)
def test_finalPositionOfSnake(n, commands, wanted):
    assert Solution().finalPositionOfSnake(n, commands) == wanted
