import pytest
from .solution import Solution

test_cases = [
    pytest.param("121", True, id="tc1"),
    pytest.param("-121", False, id="tc2"),
    pytest.param("10", False, id="tc3"),
]


@pytest.mark.parametrize("x, wanted", test_cases)
def test_isPalindrome(x, wanted):
    assert Solution().isPalindrome(x) == wanted
