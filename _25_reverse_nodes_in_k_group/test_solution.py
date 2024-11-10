import pytest
from leetgo_py import *
from .solution import Solution

test_cases = [
    pytest.param(
        deserialize("ListNode", "[1,2,3,4,5]"),
        2,
        deserialize("ListNode", "[2,1,4,3,5]"),
        id="tc1",
    ),
    pytest.param(
        deserialize("ListNode", "[1,2,3,4,5]"),
        3,
        deserialize("ListNode", "[3,2,1,4,5]"),
        id="tc3",
    ),
]


@pytest.mark.parametrize("head, k, wanted", test_cases)
def test_reverseKGroup(head, k, wanted):
    actual = Solution().reverseKGroup(head, k)
    assert serialize(actual, "ListNode") == serialize(wanted, "ListNode")
