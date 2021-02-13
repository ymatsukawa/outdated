import unittest
from test.unittest.test_tools.hashabledict import HashableDict as hd
from unittest.mock import patch
from project.tools.network.rssclient import RssClient

class TestRssClient(unittest.TestCase):
  _rss_client = None
  _feed_uris  = ['http://www.example.com']

  @classmethod
  def setUpClass(cls):
    cls._rss_client = RssClient(cls._feed_uris)

  def test_check_feeds_and_get_feed(self):
    patcher = patch('feedparser.parse')
    mock_parsed = patcher.start()
    mock_parsed.return_value = hd({'status' : 200, 'entries' : ['article_0', 'article_1'] })

    self._rss_client.check_feeds()
    self.assertEqual(len(self._rss_client.feeds), 2)

  def test_check_feeds_and_no_feed(self):
    patcher = patch('feedparser.parse')
    mock_parsed = patcher.start()
    mock_parsed.return_value = hd({'status' : 500, 'entries' : [] })
