import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 8, 6, 2, 5, 4, 8, 3, 7], 49, id="tc1"),
    pytest.param([1, 1], 1, id="tc2"),
]


@pytest.mark.parametrize("height, wanted", test_cases)
def test_maxArea(height, wanted):
    assert Solution().maxArea(height) == wanted
