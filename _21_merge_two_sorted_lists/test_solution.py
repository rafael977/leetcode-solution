import pytest
from .solution import Solution
from leetgo_py import *

test_cases = [
    pytest.param(
        deserialize("ListNode", "[1,2,4]"),
        deserialize("ListNode", "[1,3,4]"),
        deserialize("ListNode", "[1,1,2,3,4,4]"),
        id="tc1",
    ),
    pytest.param(
        deserialize("ListNode", "[]"),
        deserialize("ListNode", "[]"),
        deserialize("ListNode", "[]"),
        id="tc2",
    ),
    pytest.param(
        deserialize("ListNode", "[]"),
        deserialize("ListNode", "[0]"),
        deserialize("ListNode", "[0]"),
        id="tc3",
    ),
]


@pytest.mark.parametrize("list1, list2, wanted", test_cases)
def test_mergeTwoLists(list1, list2, wanted):
    actual = Solution().mergeTwoLists(list1, list2)
    assert serialize(actual, "ListNode") == serialize(wanted, "ListNode")
