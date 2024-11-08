import pytest
from .solution import Solution
from leetgo_py import *

test_cases = [
    pytest.param(
        [
            deserialize("ListNode", "[1,4,5]"),
            deserialize("ListNode", "[1,3,4]"),
            deserialize("ListNode", "[2,6]"),
        ],
        deserialize("ListNode", "[1,1,2,3,4,4,5,6]"),
        id="tc1",
    ),
    pytest.param(
        [],
        None,
        id="tc2",
    ),
    pytest.param(
        [
            deserialize("ListNode", "[]"),
        ],
        None,
        id="tc3",
    ),
]


@pytest.mark.parametrize("lists, wanted", test_cases)
def test_mergeKLists(lists, wanted):
    actual = Solution().mergeKLists(lists)
    assert serialize(actual, "ListNode") == serialize(wanted, "ListNode")
