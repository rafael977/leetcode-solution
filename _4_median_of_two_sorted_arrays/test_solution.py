import pytest
from .solution import Solution

test_cases = [
    pytest.param([1, 3], [2], 2, id="tc1"),
    pytest.param([1, 2], [3, 4], 2.5, id="tc2"),
    pytest.param([2], [1, 3, 4], 2.5, id="tc3"),
]


@pytest.mark.parametrize("nums1, nums2, wanted", test_cases)
def test_findMedianSortedArrays(nums1, nums2, wanted):
    assert Solution().findMedianSortedArrays(nums1, nums2) == wanted
