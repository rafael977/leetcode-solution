from solution import Solution
import pytest

test_cases = [([1, 3, 1], [3, 1, 1], 5), ([4, 1, 1], [1, 1, 3], 7)]


@pytest.mark.parametrize("energyDrinkA, energyDrinkB, wanted", test_cases)
def test_maxEnergyBoost(energyDrinkA, energyDrinkB, wanted):
    assert Solution().maxEnergyBoost(energyDrinkA, energyDrinkB) == wanted
