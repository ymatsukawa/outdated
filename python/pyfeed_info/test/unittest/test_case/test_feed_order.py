import unittest
from datetime import datetime, timezone, timedelta
from test.unittest.test_tools.hashabledict import HashableDict as hd
from project.tools.machine.feedorder import FeedOrder

class TestFeedOrder(unittest.TestCase):
  _feed_order = None
  _today      = (datetime.now(timezone.utc)).strftime("%Y-%m-%dT%H:%M:%S%z")
  _3_days_ago = (datetime.now(timezone.utc) - timedelta(0, 3600 * 24 * 3)).strftime("%Y-%m-%dT%H:%M:%S%z")
  _4_days_ago = (datetime.now(timezone.utc) - timedelta(0, 3600 * 24 * 4)).strftime("%Y-%m-%dT%H:%M:%S%z")

  @classmethod
  def setUp(self):
    entries = [
      hd({'id' : '0', 'title' : 'somebody said foo', 'published' : self._today}),
      hd({'id' : '1', 'title' : 'super bar',         'published' : self._3_days_ago}),
      hd({'id' : '2', 'title' : 'whoa blah',         'published' : self._4_days_ago})
    ]
    self._feed_order = FeedOrder(set(entries))
    self._feed_order._date_types = ['published']

  def test_filter_by_keyword_(self):
    self._feed_order.filter_by_keyword('title', ['foo', 'bar'])

    result_ids = set([elem['id'] for elem in self._feed_order.feed_set])
    self.assertEqual(len(self._feed_order.feed_set), 2)
    self.assertNotIn('2', result_ids)

  def test_filter_by_days(self):
    self._feed_order.filter_by_days(3)

    result_ids = set([elem['id'] for elem in self._feed_order.feed_set])
    self.assertEqual(len(self._feed_order.feed_set), 2)
    self.assertNotIn('2', result_ids)
