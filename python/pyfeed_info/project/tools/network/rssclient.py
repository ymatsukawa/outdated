import feedparser
from project.const import TITLE, LINK, STATUS, ENTRIES, FEED_RESPONSE_KEY, FEED_SUCCESS_RESPONSE

class RssClient:
  def __init__(self, feed_uris):
    '''
    summary:
      this class checks rss feeds and stores them
    preconditions:
      feed_uris ... should be list of rss(1.0, 2.0)
      ex. ['http://www.example.com/example/0.xml', 'http://www.example.com/example/1.rdf']
    postconditions:
      after check_feeds(), feeds property returns feeds as dictionary
    '''
    self._feed_uris = feed_uris
    self._feeds = None

  def check_feeds(self):
    '''
    summary:
      checks and get latest rss feeds
    preconditions:
      class should be initialized by feed_uris
    postconditions:
      'set' feeds property is refreshed as dict
      ex. feeds is ... set([ { 'id': 0, 'title': 'example', 'link': 'http://...', ... } ])
    '''
    self._feeds = set()
    for feed_uri in self._feed_uris:
      response = feedparser.parse(feed_uri)
      if self._is_feed_exist(response):
        self._feeds.update(response[ENTRIES])
    return self

  def _is_feed_exist(self, response):
    return (
      (STATUS in response)
      and (response[STATUS] == FEED_SUCCESS_RESPONSE[STATUS])
      and (len(response[ENTRIES]) > 0)
    )

  @property
  def feeds(self):
    return self._feeds
