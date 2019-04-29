import unittest
from parameterized import parameterized

from project.tools.machine.regex import Regex

class TestOrganizer(unittest.TestCase):
  _regex = None

  @classmethod
  def setUpClass(self):
    self._regex = Regex()

  @parameterized.expand([
    ("hello world", True),
    ("a word test", True),
    ("py th on",    True),
    ("foobarbaz",   False)
  ])
  def test_or_search(self, word, match_expectation):
    result = self._regex.or_pattern(["or", "th"]).match(word)
    if(match_expectation):
      self.assertIsNotNone(result)
    else:
      self.assertIsNone(result)


