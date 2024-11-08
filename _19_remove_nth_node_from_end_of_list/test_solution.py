import pytest
from .solution import Solution
from leetgo_py import *

test_cases = [
    pytest.param(
        deserialize("ListNode", "[1, 2, 3, 4, 5]"),
        2,
        "[1,2,3,5]",
        id="tc1",
    ),
    pytest.param(deserialize("ListNode", "[1]"), 1, None, id="tc2"),
    pytest.param(deserialize("ListNode", "[1, 2]"), 1, "[1]", id="tc3"),
]


@pytest.mark.parametrize("head, n, wanted", test_cases)
def test_removeNthFromEnd(head, n, wanted):
    actual = Solution().removeNthFromEnd(head, n)
    assert (actual if actual is None else serialize(actual)) == wanted
