import typing
from leetgo_py import deserialize, serialize
import pytest
from .solution import Solution

test_cases = [
    pytest.param(
        deserialize("ListNode", "[2,4,3]"),
        deserialize("ListNode", "[5,6,4]"),
        deserialize("ListNode", "[7,0,8]"),
        id="tc1",
    ),
    pytest.param(
        deserialize("ListNode", "[0]"),
        deserialize("ListNode", "[0]"),
        deserialize("ListNode", "[0]"),
        id="tc2",
    ),
    pytest.param(
        deserialize("ListNode", "[9,9,9,9,9,9,9]"),
        deserialize("ListNode", "[9,9,9,9]"),
        deserialize("ListNode", "[8,9,9,9,0,0,0,1]"),
        id="tc3",
    ),
]


@pytest.mark.parametrize("l1, l2, wanted", test_cases)
def test_addTwoNumbers(l1, l2, wanted):
    assert serialize(Solution().addTwoNumbers(l1, l2)) == serialize(wanted)
