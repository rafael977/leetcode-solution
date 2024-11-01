from .solution import Solution
import pytest

test_cases = [
    pytest.param([1, 3, 1], [3, 1, 1], 5, id="tc1"),
    pytest.param([4, 1, 1], [1, 1, 3], 7, id="tc2"),
]


@pytest.mark.parametrize("energyDrinkA, energyDrinkB, wanted", test_cases)
def test_maxEnergyBoost(energyDrinkA, energyDrinkB, wanted):
    assert Solution().maxEnergyBoost(energyDrinkA, energyDrinkB) == wanted
