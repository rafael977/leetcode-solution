import pytest
from .solution import Solution

test_cases = [
    pytest.param(4, "1211", id="tc1"),
    pytest.param(1, "1", id="tc2"),
]


@pytest.mark.parametrize("n, wanted", test_cases)
def test_countAndSay(n, wanted):
    assert Solution().countAndSay(n) == wanted
