import pytest
from .solution import Solution

test_cases = [
    pytest.param([16, 16], 2, id="tc1"),
    pytest.param([16, 17, 18], 2, id="tc2"),
    pytest.param([20, 30, 100, 110, 120], 3, id="tc3"),
]


@pytest.mark.parametrize("ages, wanted", test_cases)
def test_numFriendRequests(ages, wanted):
    assert Solution().numFriendRequests(ages) == wanted
