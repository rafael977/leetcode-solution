import pytest
from .solution import Solution

test_cases = [pytest.param(5, True, id="tc1"), pytest.param(3, False, id="tc2")]


@pytest.mark.parametrize("c, wanted", test_cases)
def test_judgeSquareSum(c, wanted):
    assert Solution().judgeSquareSum(c) == wanted
