import pytest
from .solution import Solution

test_cases = [
    pytest.param("barfoothefoobarman", ["foo", "bar"], [0, 9], id="tc1"),
    pytest.param(
        "wordgoodgoodgoodbestword", ["word", "good", "best", "word"], [], id="tc2"
    ),
    pytest.param(
        "barfoofoobarthefoobarman", ["bar", "foo", "the"], [6, 9, 12], id="tc3"
    ),
]


@pytest.mark.parametrize("s, words, wanted", test_cases)
def test_findSubstring(s, words, wanted):
    assert Solution().findSubstring(s, words) == wanted
