import pytest
from .solution import Solution

test_cases = [
    pytest.param(13, 4, 2, id="tc1"),
    pytest.param(21, 21, 0),
    pytest.param(13, 14, -1),
]


@pytest.mark.parametrize("n, k, wanted", test_cases)
def test_minChanges(n, k, wanted):
    assert Solution().minChanges(n, k) == wanted
