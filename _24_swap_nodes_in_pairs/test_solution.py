import pytest
from leetgo_py import *
from .solution import Solution

test_cases = [
    pytest.param(
        deserialize("ListNode", "[1,2,3,4]"),
        deserialize("ListNode", "[2,1,4,3]"),
        id="tc1",
    ),
    pytest.param(
        deserialize("ListNode", "[]"),
        deserialize("ListNode", "[]"),
        id="tc2",
    ),
    pytest.param(
        deserialize("ListNode", "[1,2,3]"),
        deserialize("ListNode", "[2,1,3]"),
        id="tc3",
    ),
]


@pytest.mark.parametrize("head, wanted", test_cases)
def test_swapPairs(head, wanted):
    actual = Solution().swapPairs(head)
    assert serialize(actual, "ListNode") == serialize(wanted, "ListNode")
