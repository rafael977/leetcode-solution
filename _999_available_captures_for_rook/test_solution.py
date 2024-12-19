import pytest
from .solution import Solution

test_cases = [
    pytest.param(
        [
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", "p", ".", ".", ".", "."],
            [".", ".", ".", "R", ".", ".", ".", "p"],
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", "p", ".", ".", ".", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
        ],
        3,
        id="tc1",
    ),
    pytest.param(
        [
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", "p", "p", "p", "p", "p", ".", "."],
            [".", "p", "p", "B", "p", "p", ".", "."],
            [".", "p", "B", "R", "B", "p", ".", "."],
            [".", "p", "p", "B", "p", "p", ".", "."],
            [".", "p", "p", "p", "p", "p", ".", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
        ],
        0,
        id="tc2",
    ),
    pytest.param(
        [
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", "p", ".", ".", ".", "."],
            [".", ".", ".", "p", ".", ".", ".", "."],
            ["p", "p", ".", "R", ".", "p", "B", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
            [".", ".", ".", "B", ".", ".", ".", "."],
            [".", ".", ".", "p", ".", ".", ".", "."],
            [".", ".", ".", ".", ".", ".", ".", "."],
        ],
        3,
        id="tc3",
    ),
]


@pytest.mark.parametrize("board, wanted", test_cases)
def test_numRookCaptures(board, wanted):
    assert Solution().numRookCaptures(board) == wanted