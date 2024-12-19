import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 2, 3, 4, 5], 2, [3, 4], id="tc1"),
    pytest.param([10, 1, 10, 1, 10], 3, [1, 3], id="tc2"),
    pytest.param([10, 1, 10, 1, 10], 10, [], id="tc3"),
]


@pytest.mark.parametrize("height, threshold, wanted", test_cases)
def test_(height, threshold, wanted):
    assert Solution().stableMountains(height, threshold) == wanted
