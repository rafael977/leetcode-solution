import pytest
from .solution import Solution

test_cases = [
    pytest.param("a1", False, id="tc1"),
    pytest.param("h3", True, id="tc2"),
    pytest.param("c7", False, id="tc3"),
]


@pytest.mark.parametrize("coordinates, wanted", test_cases)
def test_squareIsWhite(coordinates, wanted):
    assert Solution().squareIsWhite(coordinates) == wanted
