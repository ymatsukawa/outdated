import unittest
from test.unittest.test_tools.hashabledict import HashableDict as hd
from project.tools.machine.dataformatter import DataFormatter

class TestOrganizer(unittest.TestCase):
  _data_formatter = None

  @classmethod
  def setUpClass(cls):
    entry_set = [
      hd ({ 'id': '0', 'title': 'hello world', 'link' : 'http://example.com/0' }),
      hd ({ 'id': '1', 'title': 'hello word',  'link' : 'http://example.com/1' }),
      hd ({ 'id': '2', 'title': 'hello wood',  'link' : 'http://example.com/2' }),
    ]
    cls._data_formatter = DataFormatter(entry_set)

  def test_data_format_to_post(self):
    self._data_formatter.format_to_post()

    self.assertIsInstance(self._data_formatter.formatted_data, list)
    self.assertEqual(len(self._data_formatter.formatted_data), 3)

    expected_first_elem = { 'title' : 'hello world', 'link' : 'http://example.com/0' }
    self.assertDictEqual(self._data_formatter.formatted_data[0], expected_first_elem)


