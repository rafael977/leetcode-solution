import pytest
from .solution import Solution

test_cases = [
    pytest.param(
        [[1, 1, 1], [1, 0, 1], [1, 1, 1]], [[0, 0, 0], [0, 0, 0], [0, 0, 0]], id="tc1"
    ),
    pytest.param(
        [[100, 200, 100], [200, 50, 200], [100, 200, 100]],
        [[137, 141, 137], [141, 138, 141], [137, 141, 137]],
        id="tc2",
    ),
]


@pytest.mark.parametrize("img, wanted", test_cases)
def test_imageSmoother(img, wanted):
    assert Solution().imageSmoother(img) == wanted
