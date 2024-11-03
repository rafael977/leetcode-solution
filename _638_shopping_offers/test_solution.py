import pytest
from .solution import Solution

test_cases = [
    pytest.param([2, 5], [[3, 0, 5], [1, 2, 10]], [3, 2], 14, id="tc1"),
    pytest.param([2, 3, 4], [[1, 1, 0, 4], [2, 2, 1, 9]], [1, 2, 1], 11, id="tc2"),
]


@pytest.mark.parametrize("price, special, needs, wanted", test_cases)
def test_shoppingOffers(price, special, needs, wanted):
    assert Solution().shoppingOffers(price, special, needs) == wanted
