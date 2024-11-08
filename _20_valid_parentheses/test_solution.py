import pytest
from .solution import Solution

test_cases = [
    pytest.param("()", True, id="tc1"),
    pytest.param("()[]{}", True, id="tc2"),
    pytest.param("(]", False, id="tc3"),
    pytest.param("([])", True, id="tc4"),
    pytest.param("[", False, id="tc5"),
]


@pytest.mark.parametrize("s, wanted", test_cases)
def test_isValid(s, wanted):
    assert Solution().isValid(s) == wanted
