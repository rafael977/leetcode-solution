import pytest
from .solution import NeighborSum

test_cases = [
    pytest.param(
        [[0, 1, 2], [3, 4, 5], [6, 7, 8]],
        [
            ("adjacentSum", 1, 6),
            ("adjacentSum", 4, 16),
            ("diagonalSum", 4, 16),
            ("diagonalSum", 8, 4),
        ],
        id="tc1",
    ),
    pytest.param(
        [[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]],
        [
            ("adjacentSum", 15, 23),
            ("diagonalSum", 9, 45),
        ],
        id="tc2",
    ),
]


@pytest.mark.parametrize("grid, ops", test_cases)
def test_neighbor_sum(grid, ops):
    obj = NeighborSum(grid)
    for o in ops:
        match o[0]:
            case "adjacentSum":
                assert obj.adjacentSum(o[1]) == o[2]
            case "diagonalSum":
                assert obj.diagonalSum(o[1]) == o[2]
